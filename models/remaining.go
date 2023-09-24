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

type RemainingIdReq struct {
	Id string `json:"id"`
}

type RemainingResp struct {
	Id           string  `json:"id"`
	BranchId     string  `json:"branch_id"`
	BranchName   string  `json:"branch_name"`
	CategoryId   string  `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Barcode      string  `json:"barcode"`
	Count        float64 `json:"count"`
	TotalPrice   float64 `json:"total_price"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetListRemainingReq struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	BranchId   string `json:"branch_id"`
	CategoryId string `json:"category_id"`
	Barcode    string `json:"barcode"`
}

type GetListRemainingResp struct {
	Remainings []RemainingResp `json:"remainings"`
	Count      int             `json:"count"`
}

type UpdateRemaining struct {
	BranchId   string  `json:"branch_id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      float64 `json:"count"`
}
