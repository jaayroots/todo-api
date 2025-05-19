package service

import (
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
)

type UserService interface {
	GetByUserID(userID uint) (*_userModel.UserRes, error)
	Update(userID uint, userUpdateReq *_userModel.UserUpdateReq) error
	Delete(userID uint) error
}
