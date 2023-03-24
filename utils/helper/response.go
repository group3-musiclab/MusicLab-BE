package helper

import (
	"musiclab-be/utils/consts"
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
	case strings.Contains(msg, consts.QUERY_NotFound):
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "validate"):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.AUTH_ErrorCreateToken):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorHash):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorComparePassword):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.QUERY_ErrorInsertData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.QUERY_NoRowsAffected):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorRole):
		code = http.StatusBadRequest
	}
	return code, resp
}
