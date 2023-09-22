package storage

import "market/models"

type StorageI interface {
	Branch() BranchesI
	Category() CategoriesI
	Product() ProductsI
}

type BranchesI interface {
	Create(req models.CreateBranch) (resp string, err error)
	Get(req models.BranchIdReq) (resp models.Branch, err error)
	GetList(req models.GetListBranchReq) (resp models.GetListBranch, err error)
	Update(req models.Branch) (string, error)
	Delete(req models.BranchIdReq) (string, error)
}

type CategoriesI interface {
	Create(req models.CreateCategory) (resp string, err error)
	Get(req models.CategoryIdReq) (resp models.CategoryResp, err error)
	GetList(req models.GetListCategoryReq) (resp models.GetListCategory, err error)
	Update(req models.Category) (string, error)
	Delete(req models.CategoryIdReq) (string, error)
}

type ProductsI interface {
	Create(req models.CreateProduct) (resp string, err error)
	Get(req models.ProductIdReq) (resp models.ProductResp, err error)
	GetList(req models.GetListProductReq) (resp models.GetListProductResp, err error)
	Update(req models.Product) (string, error)
	Delete(req models.ProductIdReq) (string, error)
}
