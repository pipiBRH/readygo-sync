package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pipiBRH/readygo-sync/config"
)

func main() {
	config, err := config.NewConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(config.DataBase["mysql"].Driver, config.DataBase["mysql"].GetConnectionURL())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
