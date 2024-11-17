package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicate = errors.New("账号冲突")
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	// 存毫秒
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflicts uint16 = 1062
		if mysqlErr.Number == uniqueConflicts {
			// 邮箱索引
			return ErrDuplicate
		}
	}
	return err
}

type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 唯一用户
	Email    string `gorm:"unique"`
	Password string
	Ctime    int64
	Utime    int64
}
