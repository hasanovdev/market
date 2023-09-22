package db

import (
	"context"
	"database/sql"
	"fmt"
	"market/models"
	"market/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{db: db}
}

func (c *categoryRepo) Create(req models.CreateCategory) (resp string, err error) {
	id := uuid.NewString()

	query := `
	INSERT INTO categories(
		id,
		name,
		parent_id
	) VALUES($1,$2,$3)
	`

	_, err = c.db.Exec(context.Background(), query,
		id,
		req.Name,
		helper.NewNullString(req.ParentId),
	)

	if err != nil {
		return
	}

	return id, nil
}

func (c *categoryRepo) Get(req models.CategoryIdReq) (resp models.CategoryResp, err error) {
	var parentName sql.NullString
	query := `
	SELECT 
		c1.id,
		c1.name, 
		c2.name::TEXT AS parent_name,
		c1.created_at::TEXT,
		c1.updated_at::TEXT
    FROM categories c1
    LEFT JOIN categories c2 ON c1.parent_id = c2.id
    where c1.id=$1;
	`

	var category models.CategoryResp
	if err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&parentName,
		&category.CreatedAt,
		&category.UpdatedAt,
	); err != nil {
		return
	}

	category.ParentName = parentName.String
	return category, nil
}

func (c *categoryRepo) GetList(req models.GetListCategoryReq) (resp models.GetListCategory, err error) {
	var (
		filter     = " WHERE true "
		offsetQ    = " OFFSET 0;"
		limit      = " LIMIT 10 "
		offset     = (req.Page - 1) * req.Limit
		count      int
		parentName sql.NullString
	)

	s := `
	SELECT 
		c1.id,
		c1.name, 
		c2.name::TEXT AS parent_name,
		c1.created_at::TEXT,
		c1.updated_at::TEXT
    FROM categories c1
    LEFT JOIN categories c2 ON c1.parent_id = c2.id `

	if req.Name != "" {
		req.Name = helper.ReplaceStringForSql(req.Name)
		filter += ` AND c1.name ILIKE ` + "'%" + req.Name + "%' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d ", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM categories c1` + filter

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
		var category models.CategoryResp
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&parentName,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}

		category.ParentName = parentName.String

		resp.Categories = append(resp.Categories, category)
		resp.Count = count
	}

	return resp, nil
}

func (c *categoryRepo) Update(req models.Category) (string, error) {
	query := `
	UPDATE categories SET 
		name=$2,
		parent_id=$3,
		updated_at=NOW() 
	WHERE id=$1;`

	resp, err := c.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.ParentId,
	)

	if err != nil {
		return "", err
	}

	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (c *categoryRepo) Delete(req models.CategoryIdReq) (string, error) {
	query := `DELETE FROM categories WHERE id=$1;`

	resp, err := c.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "OK", nil
}
