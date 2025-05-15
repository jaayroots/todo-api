package mapper

import (
	"fmt"

	"github.com/jaayroots/todo-api/entities"
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
	"golang.org/x/crypto/bcrypt"
)

func ToUserEntity(userReq *_userModel.UserReq) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		Email:     userReq.Email,
		Password:  string(hashedPassword),
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Avatar:    userReq.Avatar,
	}

	return userEntity, nil
}

func ToUserUpdateEntity(userReq *_userModel.UserUpdateReq) (*entities.User, error) {

	userEntity := &entities.User{
		Email:     userReq.Email,
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Avatar:    userReq.Avatar,
	}

	return userEntity, nil
}

func ToUserRes(user *entities.User) *_userModel.UserRes {

	fmt.Println(user.Email)
	return &_userModel.UserRes{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
	}
}
