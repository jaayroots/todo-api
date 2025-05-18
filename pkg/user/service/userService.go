package service

import (
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
)

type UserService interface {
	GetByUserID(userID int) (*_userModel.UserRes, error)
	Update(userID int, userUpdateReq *_userModel.UserUpdateReq) error
	Delete(userID int) error
}
