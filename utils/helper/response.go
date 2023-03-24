package helper

import (
	"net/http"
	"strings"
)

func Response(message string) map[string]any {
	return map[string]any{
		"message": message,
	}
}

func ResponseWithData(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}

func ErrorResponse(err error) (int, interface{}) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError
	msg := err.Error()

	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "Atoi"):
		resp["message"] = "id must be number, cannot be string"
		code = http.StatusNotFound
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
	case strings.Contains(msg, "not found"):
		resp["message"] = "data not found"
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "hashedPassword"):
		resp["message"] = "password do not match"
		code = http.StatusForbidden
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	}
	return code, resp
}
