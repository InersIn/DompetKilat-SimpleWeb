package models

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}