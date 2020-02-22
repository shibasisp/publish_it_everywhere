package config

var (
	// Env variable holds the environmental value of `ENV`
	Env = "dev"
	// Port variable holds the environmental value of `PORT`
	Port = "8080"
	// SentryDSN ...
	SentryDSN = ""
)

// Initialize init all the config for the service
func Initialize(configPath ...string) {

}
