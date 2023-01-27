package request

type ProductRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

type ProductRequestUpdate struct {
	ProductRequest ProductRequest `gorm:"embedded"`
}
