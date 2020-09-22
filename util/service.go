package util

import (
	"strconv"
	"strings"
)

// MakeSuccessResponse - Send status = success
func MakeSuccessResponse() map[string]interface{} {
	response := make(map[string]interface{}, 0)
	response["status"] = "success"
	return response
}

// MakeCustomSuccessResponse - Send status = success
func MakeCustomSuccessResponse(data map[string]interface{}) map[string]interface{} {
	response := make(map[string]interface{}, 0)
	for k, v := range data {
		response[k] = v
	}
	response["status"] = "success"
	return response
}

// MakeFailureResponse - Send status = failed
func MakeFailureResponse(err error) map[string]interface{} {
	response := make(map[string]interface{}, 0)
	response["status"] = err.Error()
	return response
}

// IsEmpty - Check if a string is empty
func IsEmpty(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true
	}
	return false
}

// IsEmptyI - Check if a string is empty
func IsEmptyI(i interface{}) bool {
	if i == nil {
		return true
	}
	if len(strings.TrimSpace(i.(string))) == 0 {
		return true
	}
	return false
}

//StrToInt - Convert string to integer
func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
