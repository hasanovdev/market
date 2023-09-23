package db

import (
	"context"
	"fmt"
	"market/config"
	"market/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type store struct {
	db                  *pgxpool.Pool
	branches            *branchRepo
	categories          *categoryRepo
	products            *productRepo
	remainings          *remainingRepo
	comingTables        *comingTableRepo
	comingTableProducts *comingTableProductRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		fmt.Println("NewWithConfig:", err.Error())
		return nil, err
	}

	return &store{
		db: pool,
	}, nil
}

func (s *store) Branch() storage.BranchesI {
	if s.branches == nil {
		s.branches = NewBranchRepo(s.db)
	}

	return s.branches
}

func (s *store) Category() storage.CategoriesI {
	if s.categories == nil {
		s.categories = NewCategoryRepo(s.db)
	}

	return s.categories
}

func (s *store) Product() storage.ProductsI {
	if s.products == nil {
		s.products = NewProductRepo(s.db)
	}

	return s.products
}

func (s *store) Remaining() storage.RemainingsI {
	if s.remainings == nil {
		s.remainings = NewRemainingRepo(s.db)
	}

	return s.remainings
}

func (s *store) ComingTable() storage.ComingTablesI {
	if s.comingTables == nil {
		s.comingTables = NewComingTableRepo(s.db)
	}

	return s.comingTables
}

func (s *store) ComingTableProduct() storage.ComingTableProductsI {
	if s.comingTableProducts == nil {
		s.comingTableProducts = NewComingTableProductRepo(s.db)
	}

	return s.comingTableProducts
}
