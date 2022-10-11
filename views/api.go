package views

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}
