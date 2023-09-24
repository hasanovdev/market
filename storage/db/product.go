package db

import (
	"context"
	"fmt"
	"market/models"
	"market/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{db: db}
}

func (p *productRepo) Create(req models.CreateProduct) (resp string, err error) {
	id := uuid.NewString()

	query := `
	INSERT INTO products(
		id,
		name,
		price,
		barcode,
		category_id
	) VALUES($1,$2,$3,$4,$5)
	`

	_, err = p.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)

	if err != nil {
		return
	}

	return id, nil
}

func (p *productRepo) Get(req models.ProductIdReq) (resp models.ProductResp, err error) {
	query := `
	SELECT 
		p.id,
		p.name,
		p.price,
		p.barcode,
		p.category_id,
		c.name, 
		p.created_at::TEXT,
		p.updated_at::TEXT
    FROM products p
    JOIN categories c ON p.category_id = c.id
    where p.id=$1;
	`

	var product models.ProductResp
	if err = p.db.QueryRow(context.Background(), query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.CategoryId,
		&product.CategoryName,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		return
	}

	return product, nil
}

func (p *productRepo) GetList(req models.GetListProductReq) (resp models.GetListProductResp, err error) {
	var (
		filter  = " WHERE true "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		p.id,
		p.name,
		p.price,
		p.barcode,
		p.category_id,
		c.name, 
		p.created_at::TEXT,
		p.updated_at::TEXT
    FROM products p
	JOIN categories c ON p.category_id = c.id `

	if req.Name != "" {
		req.Name = helper.ReplaceStringForSql(req.Name)
		filter += ` AND p.name ILIKE ` + "'%" + req.Name + "%' "
	}
	if req.Barcode != "" {
		filter += ` AND p.barcode='` + req.Barcode + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d ", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM products p` + filter

	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	err = p.db.QueryRow(context.Background(), countS).Scan(&count)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var product models.ProductResp
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Barcode,
			&product.CategoryId,
			&product.CategoryName,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Products = append(resp.Products, product)
		resp.Count = count
	}

	return resp, nil
}

func (p *productRepo) Update(req models.Product) (string, error) {
	query := `
	UPDATE products SET 
		name=$2,
		price=$3,
		barcode=$4,
		category_id=$5,
		updated_at=NOW() 
	WHERE id=$1;`

	resp, err := p.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)

	if err != nil {
		return "", err
	}

	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (p *productRepo) Delete(req models.ProductIdReq) (string, error) {
	query := `DELETE FROM products WHERE id=$1;`

	resp, err := p.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "OK", nil
}
