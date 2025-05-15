package server

import (
	_userController "github.com/jaayroots/todo-api/pkg/user/controller"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"
	_userService "github.com/jaayroots/todo-api/pkg/user/service"
)

func (s *echoServer) usersRouter() {
	router := s.app.Group("/user")
	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	userService := _userService.NewUserServiceImpl(userRepository)
	userController := _userController.NewUserControllerImpl(userService)

	router.GET("/:userID", userController.GetByUserID)
	router.PATCH("/:userID", userController.Update)
	router.DELETE("/:userID", userController.Delete)
}
