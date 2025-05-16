package service

import (
	_authModel "github.com/jaayroots/todo-api/pkg/auth/model"
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
)

type AuthService interface {
	Login(loginReq *_authModel.LoginReq) (*_authModel.LoginRes, error)
	Logout(userID uint64) error
	Register(userReq *_userModel.UserReq) (*_userModel.UserRes, error)
	Authorizing(token string) (*_authModel.LoginRes, error)
}
