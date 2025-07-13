package utils

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type PaginationMeta struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
	Limit        int `json:"limit"`
}

func Success(status int, message string, data interface{}, meta interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func Error(status int, message string) Response {
	return Response{
		Status:  status,
		Message: message,
	}
}
