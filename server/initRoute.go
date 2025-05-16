package server

import (
	_authController "github.com/jaayroots/todo-api/pkg/auth/controller"
	_authRepository "github.com/jaayroots/todo-api/pkg/auth/repository"
	_authService "github.com/jaayroots/todo-api/pkg/auth/service"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"
)

func (s *echoServer) initRoute() {

	authorizingMiddleware := s.getAuthorizingMiddleware()

	s.authRouter(authorizingMiddleware)
	s.healthRouter()
	s.todosRouter()
	s.usersRouter()
}

func (s *echoServer) getAuthorizingMiddleware() *authorizingMiddleware {

	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	authRepository := _authRepository.NewSessionRepositoryImpl(s.db, s.app.Logger)

	authService := _authService.NewAuthServiceImpl(userRepository, authRepository)
	authController := _authController.NewUserControllerImpl(authService)
	return &authorizingMiddleware{
		authController,
	}
}
