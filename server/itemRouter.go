package server

import (
	_itemController "github.com/jaayroots/todo-api/pkg/item/controller"
	_itemRepository "github.com/jaayroots/todo-api/pkg/item/repository"
	_itemService "github.com/jaayroots/todo-api/pkg/item/service"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"

)

func (s *echoServer) itemsRouter(m *authorizingMiddleware) {
	router := s.app.Group("/item")

	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)

	itemService := _itemService.NewItemServiceImpl(itemRepository, userRepository)
	itemController := _itemController.NewItemControllerImpl(itemService)

	router.POST("", itemController.Create, m.Authorizing)
	
	router.GET("/:itemID", itemController.Get, m.Authorizing)
	router.GET("/:lang/:itemID", itemController.Get, m.Authorizing)

}
