package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

var client = &http.Client{}

func GetHTTPResponse[T any](req *http.Request, jsonResponse *T) (status string, error error) {
	response, err := client.Do(req)
	if err != nil {
		return "fail", err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "fail", err
	}
	err = json.Unmarshal(bytes, jsonResponse)
	if err != nil {
		return "fail", err
	}
	return "success", nil

}
