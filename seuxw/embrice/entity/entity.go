package entity

type Response struct {
	Code       int64       `json:"code"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}
