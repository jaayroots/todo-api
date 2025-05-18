package service

import (
	"time"

	_authException "github.com/jaayroots/todo-api/pkg/auth/exception"
	_authMapper "github.com/jaayroots/todo-api/pkg/auth/mapper"
	_authModel "github.com/jaayroots/todo-api/pkg/auth/model"
	_authRepository "github.com/jaayroots/todo-api/pkg/auth/repository"
	_utils "github.com/jaayroots/todo-api/pkg/auth/utils"
	_userException "github.com/jaayroots/todo-api/pkg/user/exception"
	_userMapper "github.com/jaayroots/todo-api/pkg/user/mapper"
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"
)

type authServiceImpl struct {
	userRepository    _userRepository.UserRepository
	sessionRepository _authRepository.SessionRepository
}

func NewAuthServiceImpl(
	userRepository _userRepository.UserRepository,
	sessionRepository _authRepository.SessionRepository,
) AuthService {
	return &authServiceImpl{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (s *authServiceImpl) Register(userReq *_userModel.UserReq) (*_userModel.UserRes, error) {

	userEntity, err := _userMapper.ToUserEntity(userReq)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByEmail(userEntity.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, _userException.IsExistUser()
	}

	user, err = s.userRepository.Create(userEntity)
	if err != nil {
		return nil, err
	}

	userRes := _userMapper.ToUserRes(user)

	return userRes, nil
}

func (s *authServiceImpl) Login(loginReq *_authModel.LoginReq) (*_authModel.LoginRes, error) {

	user, err := s.userRepository.FindByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, _userException.NotFoundUser()
	}

	isValid := _utils.CheckPasswordHash(loginReq.Password, user.Password)
	if !isValid {
		return nil, _authException.AuthenticationFailed()
	}

	token, exp, err := _utils.HashToken(_userMapper.ToUserRes(user), 24)
	if err != nil {
		return nil, err
	}

	session := _authMapper.ToSessionEntity(user, token, exp, loginReq.IpAddress)

	err = s.Logout(session.UserID)
	if err != nil {
		return nil, err
	}

	_, err = s.sessionRepository.Create(session)
	if err != nil {
		return nil, err
	}

	return _authMapper.ToAuthRes(user, token), nil
}

func (s *authServiceImpl) Logout(userID int) error {

	user, err := s.userRepository.GetByUserID(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return _userException.NotFoundUser()
	}

	err = s.sessionRepository.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *authServiceImpl) Authorizing(token string) (*_authModel.LoginRes, error) {

	session, err := s.sessionRepository.FindByToken(token)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, _authException.TokenInvalid()
	}
	if time.Now().After(session.ExpiresAt) {
		return nil, _authException.TokenInvalid()
	}

	user, err := s.userRepository.GetByUserID(session.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, _authException.TokenInvalid()
	}

	return _authMapper.ToAuthRes(user, token), nil
}
