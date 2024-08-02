package utils

type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *ErrorMessage) Error() string {
	return e.Message
}

func (e *ErrorMessage) ErrorCode() int {
	return e.Code
}
