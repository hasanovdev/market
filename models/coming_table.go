package models

type ComingTable struct {
	Id        string `json:"id"`
	ComingId  string `json:"coming_id"`
	BranchId  string `json:"branch_id"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateComingTable struct {
	BranchId string `json:"branch_id"`
	DateTime string `json:"date_time"`
}

type ComingTableIdReq struct {
	Id string `json:"id"`
}

type ComingTableResp struct {
	Id         string `json:"id"`
	ComingId   string `json:"coming_id"`
	BranchName string `json:"branch_name"`
	DateTime   string `json:"date_time"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type GetListComingTableReq struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	ComingId string `json:"coming_id"`
	BranchId string `json:"branch_id"`
}

type GetListComingTable struct {
	ComingTables []ComingTableResp `json:"coming_tables"`
	Count        int               `json:"count"`
}

type UpdateComingTable struct {
	Id       string `json:"id"`
	BranchId string `json:"branch_id"`
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
}
