package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var initDone = false

// Common variables for all the projects.
var (
	// Port ...
	Port string
	// DatabaseURL is the URL of the database
	DatabaseURL string
	// DatabaseName is the name of the database
	DatabaseName string

	// LinkedinClientID ...
	LinkedinClientID string
	// LinkedinClientSecret ...
	LinkedinClientSecret string
	// SelfURL ...
	SelfURL string
)

// RuntimeValue contains data for a single value which had to be fetched
// from the environment during runtime.
type RuntimeValue struct {
	Ptr     *string
	Default string
}

// Vars contains all the maps of env key to the pointer memory
var Vars = map[string]RuntimeValue{
	"PORT":                   {&Port, "8888"},
	"DBURL":                  {&DatabaseURL, "mongodb://localhost:27017"},
	"DBNAME":                 {&DatabaseName, "mattermost-hackathon"},
	"LINKEDIN_CLIENT_ID":     {&LinkedinClientID, ""},
	"LINKEDIN_CLIENT_SECRET": {&LinkedinClientSecret, ""},
	"SELF_URL":               {&SelfURL, "http://localhost:8888"},
}

var cfg = make(map[string]string)

// AddNewEnvEntry is used to add a new env entry into the Vars variable
//   envKey - What key this variable is defined as in env
//   def    - What should be the default value if it is not defined
//   ptr    - Pointer to the variable
func AddNewEnvEntry(envkey, def string, ptr *string) {
	if initDone {
		panic("Initialization done already")
	}
	Vars[envkey] = RuntimeValue{
		Default: def,
		Ptr:     ptr,
	}
}

// loadFromJSON is used to load the config key value pairs from a json file
func loadFromJSON(files ...string) {
	if initDone {
		panic("Initialization done already")
	}
	for _, file := range files {
		bs, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[WARN] unable to load from config file: %s\n", err)
			return
		}

		err = json.Unmarshal(bs, &cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[WARN] unable to load from config file: %s\n", err)
			return
		}
	}
}

// Initialize initializes all the variables for this package.
// This function is always used after setting the source and other
// required variables to be read (AddNewEnvEntry and LoadFromJSON).
//
// Example
//     env.LoadFromJSON("config.json", "dash.json")
//     env.AddEnvEntry("HOST", "localhost", &host)
//     env.Initialize()
func Initialize(files ...string) {
	if initDone {
		panic("Initialization done already")
	}

	loadFromJSON(files...)

	for k, v := range Vars {
		mustEnv(k, v.Ptr, v.Default)
	}

	initDone = true
}

// mustEnv get the env variable with the name 'key' and store it in 'value'
func mustEnv(key string, value *string, defaultVal string) {
	val, ok := cfg[key]
	if ok && val != "" {
		*value = val
		return
	}

	if *value = os.Getenv(key); *value == "" {
		*value = defaultVal
		fmt.Printf("%s env variable not set, using default value.\n", key)
	}
}
