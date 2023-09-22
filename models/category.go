package models

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ParentId  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryResp struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ParentName string `json:"parent_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateCategory struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type CategoryIdReq struct {
	Id string `json:"id"`
}

type GetListCategoryReq struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Name  string `json:"name"`
}

type GetListCategory struct {
	Categories []CategoryResp `json:"categories"`
	Count      int            `json:"count"`
}

type UpdateCategory struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}
