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
	case strings.Contains(msg, consts.AWS_ErrorUpload):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.QUERY_ErrorUpdateData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorEmptyPassword):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.QUERY_ErrorDeleteData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.QUERY_ErrorReadData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "schedule not available"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "minimum start date input is today"):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.CHAT_ErrorMentorID):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.CHAT_ErrorStudentID):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.AUTH_DuplicateEmail):
		code = http.StatusBadRequest
	case strings.Contains(msg, "token oauth cannot empty"):
		code = http.StatusBadRequest
	}
	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "Unauthorized") {
		code = http.StatusUnauthorized
	} else if strings.Contains(msg, "not found") {
		code = http.StatusNotFound
	}

	return code, resp
}
