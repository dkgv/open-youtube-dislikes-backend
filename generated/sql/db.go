// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.deleteDislikeStmt, err = db.PrepareContext(ctx, deleteDislike); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteDislike: %w", err)
	}
	if q.findAggregateDislikeByIDStmt, err = db.PrepareContext(ctx, findAggregateDislikeByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindAggregateDislikeByID: %w", err)
	}
	if q.findDislikeStmt, err = db.PrepareContext(ctx, findDislike); err != nil {
		return nil, fmt.Errorf("error preparing query FindDislike: %w", err)
	}
	if q.findNVideosByIDHashStmt, err = db.PrepareContext(ctx, findNVideosByIDHash); err != nil {
		return nil, fmt.Errorf("error preparing query FindNVideosByIDHash: %w", err)
	}
	if q.findUserByIDStmt, err = db.PrepareContext(ctx, findUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByID: %w", err)
	}
	if q.findVideoDetailsByIDStmt, err = db.PrepareContext(ctx, findVideoDetailsByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindVideoDetailsByID: %w", err)
	}
	if q.getDislikeCountStmt, err = db.PrepareContext(ctx, getDislikeCount); err != nil {
		return nil, fmt.Errorf("error preparing query GetDislikeCount: %w", err)
	}
	if q.insertAggregateDislikeStmt, err = db.PrepareContext(ctx, insertAggregateDislike); err != nil {
		return nil, fmt.Errorf("error preparing query InsertAggregateDislike: %w", err)
	}
	if q.insertDislikeStmt, err = db.PrepareContext(ctx, insertDislike); err != nil {
		return nil, fmt.Errorf("error preparing query InsertDislike: %w", err)
	}
	if q.insertUserStmt, err = db.PrepareContext(ctx, insertUser); err != nil {
		return nil, fmt.Errorf("error preparing query InsertUser: %w", err)
	}
	if q.updateAggregateDislikeStmt, err = db.PrepareContext(ctx, updateAggregateDislike); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAggregateDislike: %w", err)
	}
	if q.upsertVideoDetailsStmt, err = db.PrepareContext(ctx, upsertVideoDetails); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertVideoDetails: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteDislikeStmt != nil {
		if cerr := q.deleteDislikeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteDislikeStmt: %w", cerr)
		}
	}
	if q.findAggregateDislikeByIDStmt != nil {
		if cerr := q.findAggregateDislikeByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAggregateDislikeByIDStmt: %w", cerr)
		}
	}
	if q.findDislikeStmt != nil {
		if cerr := q.findDislikeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findDislikeStmt: %w", cerr)
		}
	}
	if q.findNVideosByIDHashStmt != nil {
		if cerr := q.findNVideosByIDHashStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findNVideosByIDHashStmt: %w", cerr)
		}
	}
	if q.findUserByIDStmt != nil {
		if cerr := q.findUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByIDStmt: %w", cerr)
		}
	}
	if q.findVideoDetailsByIDStmt != nil {
		if cerr := q.findVideoDetailsByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findVideoDetailsByIDStmt: %w", cerr)
		}
	}
	if q.getDislikeCountStmt != nil {
		if cerr := q.getDislikeCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDislikeCountStmt: %w", cerr)
		}
	}
	if q.insertAggregateDislikeStmt != nil {
		if cerr := q.insertAggregateDislikeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertAggregateDislikeStmt: %w", cerr)
		}
	}
	if q.insertDislikeStmt != nil {
		if cerr := q.insertDislikeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertDislikeStmt: %w", cerr)
		}
	}
	if q.insertUserStmt != nil {
		if cerr := q.insertUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertUserStmt: %w", cerr)
		}
	}
	if q.updateAggregateDislikeStmt != nil {
		if cerr := q.updateAggregateDislikeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAggregateDislikeStmt: %w", cerr)
		}
	}
	if q.upsertVideoDetailsStmt != nil {
		if cerr := q.upsertVideoDetailsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertVideoDetailsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	deleteDislikeStmt            *sql.Stmt
	findAggregateDislikeByIDStmt *sql.Stmt
	findDislikeStmt              *sql.Stmt
	findNVideosByIDHashStmt      *sql.Stmt
	findUserByIDStmt             *sql.Stmt
	findVideoDetailsByIDStmt     *sql.Stmt
	getDislikeCountStmt          *sql.Stmt
	insertAggregateDislikeStmt   *sql.Stmt
	insertDislikeStmt            *sql.Stmt
	insertUserStmt               *sql.Stmt
	updateAggregateDislikeStmt   *sql.Stmt
	upsertVideoDetailsStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		deleteDislikeStmt:            q.deleteDislikeStmt,
		findAggregateDislikeByIDStmt: q.findAggregateDislikeByIDStmt,
		findDislikeStmt:              q.findDislikeStmt,
		findNVideosByIDHashStmt:      q.findNVideosByIDHashStmt,
		findUserByIDStmt:             q.findUserByIDStmt,
		findVideoDetailsByIDStmt:     q.findVideoDetailsByIDStmt,
		getDislikeCountStmt:          q.getDislikeCountStmt,
		insertAggregateDislikeStmt:   q.insertAggregateDislikeStmt,
		insertDislikeStmt:            q.insertDislikeStmt,
		insertUserStmt:               q.insertUserStmt,
		updateAggregateDislikeStmt:   q.updateAggregateDislikeStmt,
		upsertVideoDetailsStmt:       q.upsertVideoDetailsStmt,
	}
}
