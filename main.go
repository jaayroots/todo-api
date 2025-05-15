package main

import (
	"github.com/jaayroots/todo-api/config"
	"github.com/jaayroots/todo-api/database"
	"github.com/jaayroots/todo-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
