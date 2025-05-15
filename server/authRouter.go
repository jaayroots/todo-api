package server

import (
	_userController "github.com/jaayroots/todo-api/pkg/user/controller"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"
	_userService "github.com/jaayroots/todo-api/pkg/user/service"
)

func (s *echoServer) authRouter() {
	router := s.app.Group("/auth")

	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	userService := _userService.NewUserServiceImpl(userRepository)
	userController := _userController.NewUserControllerImpl(userService)

	router.POST("/register", userController.Create)
}
