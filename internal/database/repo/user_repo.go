package repo

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type UserRepo struct {
	querier db.Querier
}

func NewUserRepo(conn *sql.DB) *UserRepo {
	return &UserRepo{querier: db.New(conn)}
}

func (v *UserRepo) FindByID(ctx context.Context, id string) (db.User, error) {
	return v.querier.FindUserByID(ctx, id)
}

func (v *UserRepo) Insert(ctx context.Context, id string) error {
	return v.querier.InsertUser(ctx, id)
}
