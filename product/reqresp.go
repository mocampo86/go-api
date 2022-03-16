package product

type ProductIdRequest struct {
	ProductId int64 `json:"id"`
}

type ProductFilterRequest struct {
	ProductId   int64  `json:"id"`
	ProductName string `json:"name"`
}

type ActionResultProductResponse struct {
	Success bool `json:"success"`
}
