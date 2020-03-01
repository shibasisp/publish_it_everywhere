package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"publish_it_everywhere/types"
)

// PostJSON performs http post request
func PostJSON(client *http.Client, url string, requestBody types.JSON, response interface{}) error {

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, response); err != nil {
		return err
	}

	return nil
}
