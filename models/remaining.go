package models

type Remaining struct {
	Id         string  `json:"id"`
	BranchId   string  `json:"branch_id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      float64 `json:"count"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type CreateRemaining struct {
	BranchId   string  `json:"branch_id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      float64 `json:"count"`
}
