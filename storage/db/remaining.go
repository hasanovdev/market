package db

import (
	"context"
	"market/models"

	"github.com/google/uuid"
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
