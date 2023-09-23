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

type comingTableRepo struct {
	db *pgxpool.Pool
}

func NewComingTableRepo(db *pgxpool.Pool) *comingTableRepo {
	return &comingTableRepo{db: db}
}

func (c *comingTableRepo) Create(req models.CreateComingTable) (resp string, err error) {
	id := uuid.NewString()
	comingId := helper.NewCustomIDGenerator().GenerateID()

	query := `
	INSERT INTO coming_tables(
		id,
		coming_id,
		branch_id,
		date_time
	) VALUES($1,$2,$3,$4)
	`

	_, err = c.db.Exec(context.Background(), query,
		id,
		comingId,
		req.BranchId,
		req.DateTime,
	)

	if err != nil {
		return
	}

	return id, nil
}

func (c *comingTableRepo) Get(req models.ComingTableIdReq) (resp models.ComingTableResp, err error) {
	query := `
	SELECT 
        ct.id,
        ct.coming_id,
        br.name,
        ct.date_time::TEXT,
        ct.status,
        ct.created_at::TEXT,
        ct.updated_at::TEXT
        FROM coming_tables ct
        JOIN branches br ON ct.branch_id=br.id 
    WHERE ct.id=$1;
	`

	var comingTable models.ComingTableResp
	if err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&comingTable.Id,
		&comingTable.ComingId,
		&comingTable.BranchName,
		&comingTable.DateTime,
		&comingTable.Status,
		&comingTable.CreatedAt,
		&comingTable.UpdatedAt,
	); err != nil {
		return
	}

	return comingTable, nil
}

func (c *comingTableRepo) GetList(req models.GetListComingTableReq) (resp models.GetListComingTable, err error) {
	var (
		filter  = " WHERE true "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		ct.id,
		ct.coming_id,
		br.name,
		ct.date_time::TEXT,
		ct.status,
		ct.created_at::TEXT,
		ct.updated_at::TEXT
	FROM coming_tables ct
	JOIN branches br ON ct.branch_id=br.id `

	if req.BranchId != "" {
		filter += ` AND ct.branch_id='` + req.BranchId + "' "
	}
	if req.ComingId != "" {
		filter += ` AND ct.coming_id='` + req.ComingId + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d ", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM coming_tables ct` + filter

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
		var comingTable models.ComingTableResp
		err := rows.Scan(
			&comingTable.Id,
			&comingTable.ComingId,
			&comingTable.BranchName,
			&comingTable.DateTime,
			&comingTable.Status,
			&comingTable.CreatedAt,
			&comingTable.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}
		resp.ComingTables = append(resp.ComingTables, comingTable)
		resp.Count = count
	}

	return resp, nil
}

func (c *comingTableRepo) Update(req models.ComingTable) (string, error) {
	query := `
	UPDATE coming_tables SET 
		branch_id=$2,
		date_time=$3,
		status=$4,
		updated_at=NOW() 
	WHERE id=$1;`

	resp, err := c.db.Exec(context.Background(), query,
		req.Id,
		req.BranchId,
		req.DateTime,
		req.Status,
	)

	if err != nil {
		return "", err
	}

	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (c *comingTableRepo) Delete(req models.ComingTableIdReq) (string, error) {
	query := "DELETE FROM coming_tables WHERE id=$1;"

	resp, err := c.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "OK", nil
}
