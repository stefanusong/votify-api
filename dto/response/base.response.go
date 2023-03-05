package response

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}

func New(success bool, message string, data, err any) *BaseResponse {
	return &BaseResponse{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	}
}
