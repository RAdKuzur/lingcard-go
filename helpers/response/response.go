package response

type FailedResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func CreateFailedResponse(message string) FailedResponse {
	return FailedResponse{
		Success: false,
		Message: message,
	}
}

func CreateSuccessResponse() SuccessResponse {
	return SuccessResponse{
		Success: true,
	}
}
