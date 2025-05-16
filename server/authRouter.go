package server

import (
	_authController "github.com/jaayroots/todo-api/pkg/auth/controller"
	_authService "github.com/jaayroots/todo-api/pkg/auth/service"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"
	_authRepository "github.com/jaayroots/todo-api/pkg/auth/repository"
)

func (s *echoServer) authRouter(m *authorizingMiddleware) {
	router := s.app.Group("/auth")

	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	sessionRepository := _authRepository.NewSessionRepositoryImpl(s.db, s.app.Logger)

	authService := _authService.NewAuthServiceImpl(userRepository, sessionRepository)
	authController := _authController.NewAuthControllerImpl(authService)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout, m.Authorizing)
}
