package main

import (
	"log"
	"os"

	"github.com/pipiBRH/readygo-sync/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pipiBRH/readygo-sync/config"
)

func main() {
	config, err := config.NewConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	pqlDB, err := gorm.Open(config.DataBase["postgresql"].Driver, config.DataBase["postgresql"].GetConnectionDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pqlDB.Close()

	mysqlDB, err := gorm.Open(config.DataBase["mysql"].Driver, config.DataBase["mysql"].GetConnectionDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlDB.Close()
	mysqlDB = mysqlDB.Debug()

	cards := []model.ACards{}
	pqlDB.Find(&cards)

	for _, v := range cards {
		bank := new(model.Banks)
		mysqlDB.Where("bank_code = ?", v.BankID).First(bank)
		data := model.Cards{
			BankID:       bank.ID,
			CardName:     v.Name,
			CardNickname: v.NickName,
		}
		mysqlDB.Create(&data)
	}

}
