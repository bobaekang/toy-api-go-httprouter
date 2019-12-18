package main

type JSONResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JSONErrorResponse struct {
	Error *APIError `json:"error"`
}

type APIError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}
