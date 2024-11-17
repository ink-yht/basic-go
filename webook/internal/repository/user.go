package repository

import (
	"context"
	"github.com/ink-yht/basic-go/webook/internal/domain"
	"github.com/ink-yht/basic-go/webook/internal/repository/dao"
)

var ErrDuplicate = dao.ErrDuplicate

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
