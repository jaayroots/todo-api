package service

import (
	_userException "github.com/jaayroots/todo-api/pkg/user/exception"
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"
	_userMapper "github.com/jaayroots/todo-api/pkg/user/mapper"
)

type userServiceImpl struct {
	userRepository _userRepository.UserRepository
}

func NewUserServiceImpl(
	userRepository _userRepository.UserRepository,
) UserService {
	return &userServiceImpl{userRepository}
}

func (s *userServiceImpl) GetByUserID(userID uint64) (*_userModel.UserRes, error) {

	user, err := s.userRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, _userException.NotFoundUser()
	}

	userRes := _userMapper.ToUserRes(user)
	return userRes, nil
}

func (s *userServiceImpl) Update(userID uint64, userUpdateReq *_userModel.UserUpdateReq) error {
	userEntity, err := _userMapper.ToUserUpdateEntity(userUpdateReq)
	if err != nil {
		return err
	}

	user, err := s.userRepository.GetByUserID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return _userException.NotFoundUser()
	}

	_, err = s.userRepository.Update(userID, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Delete(userID uint64) error {

	user, err := s.userRepository.GetByUserID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return _userException.NotFoundUser()
	}

	err = s.userRepository.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
