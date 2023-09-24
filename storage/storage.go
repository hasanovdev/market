package storage

import "market/models"

type StorageI interface {
	Branch() BranchesI
	Category() CategoriesI
	Product() ProductsI
	Remaining() RemainingsI
	ComingTable() ComingTablesI
	ComingTableProduct() ComingTableProductsI
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
	GetCategoryId(barcode string) (resp string, err error)
}

type ComingTablesI interface {
	Create(req models.CreateComingTable) (resp string, err error)
	Get(req models.ComingTableIdReq) (resp models.ComingTableResp, err error)
	GetList(req models.GetListComingTableReq) (resp models.GetListComingTable, err error)
	Update(req models.ComingTable) (string, error)
	Delete(req models.ComingTableIdReq) (string, error)
}

type ComingTableProductsI interface {
	Create(req models.CreateComingTableProduct) (resp string, err error)
	Get(req models.ComingTableProductIdReq) (resp models.ComingTableProductResp, err error)
	GetList(req models.GetListComingTableProductReq) (resp models.GetListComingTableProduct, err error)
	Update(req models.ComingTableProduct) (string, error)
	Delete(req models.ComingTableProductIdReq) (string, error)
}

type RemainingsI interface {
	Create(req models.CreateRemaining) (resp string, err error)
	Get(req models.RemainingIdReq) (resp models.RemainingResp, err error)
	GetList(req models.GetListRemainingReq) (resp models.GetListRemainingResp, err error)
	Update(req models.Remaining) (string, error)
	Delete(req models.RemainingIdReq) (string, error)
}
