package dao

import (
	"database/sql"
	"fmt"
	"go-tutorial/w2-error/code"
	"go-tutorial/w2-error/model"

	"github.com/pkg/errors"
)

var Db *sql.DB

type UserDao struct{}

func NewUser() *UserDao {
	return &UserDao{}
}

func (u *UserDao) GetByID(id int64) (*model.User, error) {
	user, err := getByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(code.ErrNotFound, fmt.Sprintf("dao：user_GetByID(%d) failed", id))
	}
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("dao：user_GetByID(%d) failed", id))
	}
	return user, nil
}

// 底层方法将error抛给上层进行处理，调用方决定处理逻辑
func getByID(id int64) (*model.User, error) {
	user := &model.User{}
	row := Db.QueryRow("select id ,name from user where id = ?", id)
	err := row.Scan(user.ID, user.Name)
	return user, err
}
