package main

import (
	"example.com/gosql/db"
	"example.com/gosql/films"
	"example.com/gosql/server"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&films.Film{})

	server.Init()
}
