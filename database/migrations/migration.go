package main

import (
	"github.com/jaayroots/todo-api/config"
	"github.com/jaayroots/todo-api/database"
	"github.com/jaayroots/todo-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	userMigration(tx)
	sessionMigration(tx)
	todoMigration(tx)
	itemMigration(tx)
	itemTranslationMigration(tx)

	tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
		panic(tx.Error)
	}
}

func userMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.User{})
}

func sessionMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Session{})
}

func todoMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Todo{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func itemTranslationMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.ItemTranslation{})
}
