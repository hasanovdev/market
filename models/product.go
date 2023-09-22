package models

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	CategoryId string  `json:"category_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	CategoryId string  `json:"category_id"`
}

type ProductIdReq struct {
	Id string `json:"id"`
}

type ProductResp struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Barcode      string  `json:"barcode"`
	CategoryName string  `json:"category_name"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetListProductReq struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
}

type GetListProductResp struct {
	Products []ProductResp `json:"products"`
	Count    int           `json:"count"`
}

type UpdateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	CategoryId string  `json:"category_id"`
}
