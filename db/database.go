package db

import (
	"fmt"

	"example.com/gosql/creds"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(
		mysql.Open(fmt.Sprintf("root:%s@tcp(localhost:3306)/sakila?parseTime=true", creds.SQLroot)),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
}
