package db

import (
	"context"
	"fmt"
	"market/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type remainingRepo struct {
	db *pgxpool.Pool
}

func NewRemainingRepo(db *pgxpool.Pool) *remainingRepo {
	return &remainingRepo{db: db}
}

func (r *remainingRepo) Create(req models.CreateRemaining) (resp string, err error) {
	id := uuid.NewString()
	totalPrice := req.Price * req.Count

	query := `
	INSERT INTO remainings(
		id,
		branch_id,
		category_id,
		name,
		price,
		barcode,
		count,
		total_price
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8);
	`

	_, err = r.db.Exec(context.Background(), query,
		id,
		req.BranchId,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		totalPrice,
	)

	if err != nil {
		return
	}

	return id, nil
}

func (r *remainingRepo) Get(req models.RemainingIdReq) (resp models.RemainingResp, err error) {
	query := `
	SELECT 
		r.id,
		br.name,
		ct.name,
		r.name,
		r.price,
		r.barcode,
		r.count, 
		r.total_price,
		r.created_at::TEXT,
		r.updated_at::TEXT
    FROM remainings r
    JOIN branches br ON r.branch_id = br.id
	JOIN categories ct ON r.category_id=ct.id
    where r.id=$1;
	`

	var remaining models.RemainingResp
	if err = r.db.QueryRow(context.Background(), query, req.Id).Scan(
		&remaining.Id,
		&remaining.BranchName,
		&remaining.CategoryName,
		&remaining.Name,
		&remaining.Price,
		&remaining.Barcode,
		&remaining.Count,
		&remaining.TotalPrice,
		&remaining.CreatedAt,
		&remaining.UpdatedAt,
	); err != nil {
		return
	}

	return remaining, nil
}

func (r *remainingRepo) GetList(req models.GetListRemainingReq) (resp models.GetListRemainingResp, err error) {
	var (
		filter  = " WHERE true "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		r.id,
		r.branch_id,
		br.name,
		r.category_id,
		ct.name,
		r.name,
		r.price,
		r.barcode,
		r.count, 
		r.total_price,
		r.created_at::TEXT,
		r.updated_at::TEXT
	FROM remainings r
	JOIN branches br ON r.branch_id = br.id
	JOIN categories ct ON r.category_id=ct.id `

	if req.Barcode != "" {
		filter += ` AND r.barcode='` + req.Barcode + "' "
	}
	if req.BranchId != "" {
		filter += ` AND r.branch_id='` + req.BranchId + "' "
	}
	if req.CategoryId != "" {
		filter += ` AND r.category_id='` + req.CategoryId + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d ", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM remainings r` + filter

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	err = r.db.QueryRow(context.Background(), countS).Scan(&count)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var remaining models.RemainingResp
		err := rows.Scan(
			&remaining.Id,
			&remaining.BranchId,
			&remaining.BranchName,
			&remaining.CategoryId,
			&remaining.CategoryName,
			&remaining.Name,
			&remaining.Price,
			&remaining.Barcode,
			&remaining.Count,
			&remaining.TotalPrice,
			&remaining.CreatedAt,
			&remaining.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Remainings = append(resp.Remainings, remaining)
		resp.Count = count
	}

	return resp, nil
}

func (r *remainingRepo) Update(req models.Remaining) (string, error) {
	query := `
	UPDATE remainings SET 
		branch_id=$2,
		category_id=$3,
		name=$4,
		price=$5,
		barcode=$6,
		count=$7,
		total_price=$8,
		updated_at=NOW() 
	WHERE id=$1;`

	resp, err := r.db.Exec(context.Background(), query,
		req.Id,
		req.BranchId,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.Price*req.Count,
	)

	if err != nil {
		return "", err
	}

	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (r *remainingRepo) Delete(req models.RemainingIdReq) (string, error) {
	query := `DELETE FROM remainings WHERE id=$1;`

	resp, err := r.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "OK", nil
}

func (r *remainingRepo) CheckProductExists(barcode string) (bool, string, float64) {
	var (
		count    int
		id       string
		remCount float64
	)
	query := `
	SELECT 
		id,
		count,
		COUNT(*)
	FROM remainings
	WHERE barcode=$1 GROUP BY id;
	`

	err := r.db.QueryRow(context.Background(), query, barcode).Scan(&id, &remCount, &count)
	if err != nil {
		return false, "", 0
	}

	return count > 0, id, remCount
}
