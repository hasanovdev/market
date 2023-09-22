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
