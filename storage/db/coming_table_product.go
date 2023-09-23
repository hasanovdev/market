package db

import (
	"context"
	"fmt"
	"market/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type comingTableProductRepo struct {
	db *pgxpool.Pool
}

func NewComingTableProductRepo(db *pgxpool.Pool) *comingTableProductRepo {
	return &comingTableProductRepo{db: db}
}

func (c *comingTableProductRepo) Create(req models.CreateComingTableProduct) (resp string, err error) {
	id := uuid.NewString()
	totalPrice := req.Count * req.Price

	query := `
	INSERT INTO coming_table_products(
		id,
		category_id,
		name,
		price,
		barcode,
		count,
		total_price,
		coming_table_id
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_, err = c.db.Exec(context.Background(), query,
		id,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		totalPrice,
		req.ComingTableId,
	)

	if err != nil {
		return
	}

	return id, nil
}

func (c *comingTableProductRepo) Get(req models.ComingTableProductIdReq) (resp models.ComingTableProductResp, err error) {
	query := `
	SELECT 
        cp.id,
        ct.name,
        cp.name,
        cp.price,
		cp.barcode,
		cp.count,
        cp.total_price,
		cp.coming_table_id,
        cp.created_at::TEXT,
        cp.updated_at::TEXT
        FROM coming_table_products cp
        JOIN categories ct ON cp.category_id=ct.id 
    WHERE cp.id=$1;
	`

	var comingTableProduct models.ComingTableProductResp
	if err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&comingTableProduct.Id,
		&comingTableProduct.CategoryName,
		&comingTableProduct.Name,
		&comingTableProduct.Price,
		&comingTableProduct.Barcode,
		&comingTableProduct.Count,
		&comingTableProduct.TotalPrice,
		&comingTableProduct.ComingTableId,
		&comingTableProduct.CreatedAt,
		&comingTableProduct.UpdatedAt,
	); err != nil {
		return
	}

	return comingTableProduct, nil
}

func (c *comingTableProductRepo) GetList(req models.GetListComingTableProductReq) (resp models.GetListComingTableProduct, err error) {
	var (
		filter  = " WHERE true "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		cp.id,
		ct.name,
		cp.name,
		cp.price,
		cp.barcode,
		cp.count,
		cp.total_price,
		cp.coming_table_id,
		cp.created_at::TEXT,
		cp.updated_at::TEXT
	FROM coming_table_products cp
	JOIN categories ct ON cp.category_id=ct.id `

	if req.CategoryId != "" {
		filter += ` AND cp.category_id='` + req.CategoryId + "' "
	}
	if req.Barcode != "" {
		filter += ` AND cp.barcode='` + req.Barcode + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d ", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM coming_table_products cp` + filter

	rows, err := c.db.Query(context.Background(), query)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	err = c.db.QueryRow(context.Background(), countS).Scan(&count)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var comingTableProduct models.ComingTableProductResp
		err := rows.Scan(
			&comingTableProduct.Id,
			&comingTableProduct.CategoryName,
			&comingTableProduct.Name,
			&comingTableProduct.Price,
			&comingTableProduct.Barcode,
			&comingTableProduct.Count,
			&comingTableProduct.TotalPrice,
			&comingTableProduct.ComingTableId,
			&comingTableProduct.CreatedAt,
			&comingTableProduct.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}
		resp.ComingTableProducts = append(resp.ComingTableProducts, comingTableProduct)
		resp.Count = count
	}

	return resp, nil
}

func (c *comingTableProductRepo) Update(req models.ComingTableProduct) (string, error) {
	query := `
	UPDATE coming_table_products SET 
		category_id=$2,
		name=$3,
		price=$4,
		barcode=$5,
		count=$6,
		total_price=$7,
		coming_table_id=$8,
		updated_at=NOW() 
	WHERE id=$1;`

	resp, err := c.db.Exec(context.Background(), query,
		req.Id,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.Count*req.Price,
		req.ComingTableId,
	)

	if err != nil {
		return "", err
	}

	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (c *comingTableProductRepo) Delete(req models.ComingTableProductIdReq) (string, error) {
	query := "DELETE FROM coming_table_products WHERE id=$1;"

	resp, err := c.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "OK", nil
}
