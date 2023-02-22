package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var client = &http.Client{}

func GetHTTPResponse(req *http.Request, jsonResponse interface{}) (status string, error error) {
	fmt.Println("Running GetHTTPResponse")
	response, err := client.Do(req)
	if err != nil {
		return "fail", err
	}
	defer response.Body.Close()

	fmt.Println("response", response.StatusCode)
	if response.StatusCode > 200 {
		return "fail", err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("readAll err ", err)
		return "fail", err
	}
	fmt.Println("jsonResponse ", jsonResponse)
	err = json.Unmarshal(bytes, jsonResponse)
	if err != nil {
		log.Fatal("unmarshal err ", err)
		return "fail", err
	}
	return "success", nil

}

func CreateSuccessResponse(data any) (successResponse map[string]any) {
	response := make(map[string]any)
	response["status"] = 200
	response["message"] = "success"
	response["data"] = data
	return response
}

func CreateErrorResponse(statusCode int, message string) (errorResponse map[string]any) {
	response := make(map[string]any)
	response["status"] = statusCode
	response["message"] = message
	return response
}
