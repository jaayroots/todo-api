package server

func (s *echoServer) initRoute() {
	s.authRouter()
	s.healthRouter()
	s.todosRouter()
	s.usersRouter()
}
