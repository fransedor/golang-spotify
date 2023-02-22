package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var client = &http.Client{}

func GetHTTPResponse(req *http.Request, jsonResponse interface{}) (status string, newResponse CustomResponse) {
	fmt.Println("Running GetHTTPResponse")
	response, err := client.Do(req)
	if err != nil {
		return "fail", CreateErrorResponse(500, "Request failed")
	}
	defer response.Body.Close()

	fmt.Println("response", response.StatusCode)

	if response.StatusCode == 401 {
		err = errors.New("expired token")
		return "fail", CreateErrorResponse(401, "Expired token")
	}
	if response.StatusCode > 200 {
		err = errors.New("Not OK")
		return "fail", CreateErrorResponse(500, "Not OK response")
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("readAll err ", err)
		return "fail", CreateErrorResponse(500, "Failed to read response")
	}

	err = json.Unmarshal(bytes, jsonResponse)
	if err != nil {
		fmt.Println("unmarshal err ", err)
		return "fail", CreateErrorResponse(500, "Failed to unmarshal response")
	}
	return "success", CreateSuccessResponse(nil)

}

type CustomResponse struct {
	Status  int
	Message string
	Data    any
}

func CreateSuccessResponse(data any) (successResponse CustomResponse) {
	var response CustomResponse
	response.Status = 200
	response.Message = "success"
	response.Data = data
	return response
}

func CreateErrorResponse(statusCode int, message string) (errorResponse CustomResponse) {
	var response CustomResponse
	response.Status = statusCode
	response.Message = message
	response.Data = nil
	return response
}
