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

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{db: db}
}

func (b *branchRepo) Create(req models.CreateBranch) (resp string, err error) {
	id := uuid.NewString()

	query := `
	INSERT INTO branches(
		id,
		name,
		address,
		phone_number
	) VALUES($1,$2,$3,$4)
	`

	_, err = b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *branchRepo) Get(req models.BranchIdReq) (resp models.Branch, err error) {
	query := `
	SELECT 
		id,
		name,
		address,
		phone_number,
		created_at::TEXT,
		updated_at::TEXT
	FROM branches WHERE id=$1;
	`

	var branch models.Branch
	if err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.PhoneNumber,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	); err != nil {
		return
	}

	return branch, nil
}

func (b *branchRepo) GetList(req models.GetListBranchReq) (resp models.GetListBranch, err error) {
	var (
		filter  = " WHERE true "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		id,
		name,
		address,
		phone_number,
		created_at::TEXT,
		updated_at::TEXT 
	FROM branches `

	if req.Name != "" {
		req.Name = helper.ReplaceStringForSql(req.Name)
		filter += ` AND name ILIKE ` + "'%" + req.Name + "%' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d ", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM branches` + filter

	rows, err := b.db.Query(context.Background(), query)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	err = b.db.QueryRow(context.Background(), countS).Scan(&count)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var branch models.Branch
		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.PhoneNumber,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}
		resp.Branches = append(resp.Branches, branch)
		resp.Count = count
	}

	return resp, nil
}

func (b *branchRepo) Update(req models.Branch) (string, error) {
	query := `
	UPDATE branches SET 
		name=$2,
		address=$3,
		phone_number=$4,
		updated_at=NOW() 
	WHERE id=$1;`

	resp, err := b.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (b *branchRepo) Delete(req models.BranchIdReq) (string, error) {
	query := "DELETE FROM branches WHERE id=$1;"

	resp, err := b.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "OK", nil
}
