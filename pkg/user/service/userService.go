package service

import (
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
)

type UserService interface {
	Create(userReq *_userModel.UserReq) (*_userModel.UserRes, error)
	GetByUserID(userID uint64) (*_userModel.UserRes, error)
	Update(userID uint64, userUpdateReq *_userModel.UserUpdateReq) error
	Delete(userID uint64) error
}
