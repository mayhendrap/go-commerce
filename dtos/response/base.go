package response

type DefaultResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginateResponse struct {
	DefaultResponse DefaultResponse `gorm:"embedded"`
	Pagination      interface{}     `json:"pagination"`
}
