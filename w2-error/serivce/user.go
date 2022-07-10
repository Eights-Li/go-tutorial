package service

import (
	"github.com/pkg/errors"
	"go-tutorial/w2-error/dao"
	"go-tutorial/w2-error/model"
)

type UserService struct{}

func NewUser() *UserService {
	return &UserService{}
}

func (u *UserService) GetUser(ID int64) (*model.User, error) {
	userDao := dao.NewUser()
	user, err := userDao.GetByID(ID)
	if err != nil {
		return nil, errors.WithMessagef(err, "service：user_GetUser(%d) failed", ID)
	}
	return user, err
}
