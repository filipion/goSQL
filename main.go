package main

import (
	"example.com/gosql/db"
	"example.com/gosql/models"
	"example.com/gosql/server"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.Film{})

	server.Init()
}
