package models

type ComingTableProduct struct {
	Id            string  `json:"id"`
	CategoryId    string  `json:"category_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         float64 `json:"count"`
	TotalPrice    float64 `json:"total_price"`
	ComingTableId string  `json:"coming_table_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type CreateComingTableProduct struct {
	CategoryId    string  `json:"category_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         float64 `json:"count"`
	ComingTableId string  `json:"coming_table_id"`
}

type ComingTableProductIdReq struct {
	Id string `json:"id"`
}

type ComingTableProductResp struct {
	Id            string  `json:"id"`
	CategoryName  string  `json:"category_name"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         float64 `json:"count"`
	TotalPrice    float64 `json:"total_price"`
	ComingTableId string  `json:"coming_table_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type GetListComingTableProductReq struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	CategoryId string `json:"category_id"`
	Barcode    string `json:"barcode"`
}

type GetListComingTableProduct struct {
	ComingTableProducts []ComingTableProductResp `json:"coming_table_products"`
	Count               int                      `json:"count"`
}

type UpdateComingTableProduct struct {
	CategoryId    string  `json:"category_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         float64 `json:"count"`
	ComingTableId string  `json:"coming_table_id"`
}
