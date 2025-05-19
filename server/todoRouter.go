package server

import (
	_todoController "github.com/jaayroots/todo-api/pkg/todo/controller"
	_todoRepository "github.com/jaayroots/todo-api/pkg/todo/repository"
	_todoService "github.com/jaayroots/todo-api/pkg/todo/service"
)

func (s *echoServer) todosRouter(m *authorizingMiddleware) {
	router := s.app.Group("/todo")

	todoRepository := _todoRepository.NewTodoRepositoryImpl(s.db, s.app.Logger)
	todoService := _todoService.NewTodoServiceImpl(todoRepository)
	todoController := _todoController.NewTodoControllerImpl(todoService)

	router.POST("/search", todoController.FindAll, m.Authorizing)
	router.GET("/:todoID", todoController.Get, m.Authorizing)
	router.POST("", todoController.Create, m.Authorizing)
	router.PATCH("/:todoID", todoController.Update, m.Authorizing)
	router.DELETE("/:todoID", todoController.Delete, m.Authorizing)
}
