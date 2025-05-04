package database

import (
	"client-server-api/server/database/model"
	"context"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Start() {
	Connect()
	Migrate()
}

func Connect() {

	database, err := gorm.Open(sqlite.Open("file:db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DB = database
}

func Migrate() {

	DB.AutoMigrate(&model.DollarPrice{})
}

func InsertDollar(dollarPrice *model.DollarPrice) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	defer cancel()

	DB.WithContext(ctx).Create(dollarPrice)
}
