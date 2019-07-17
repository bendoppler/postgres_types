package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/piendop/postgresql/config"
)

var (
	once sync.Once
	db   *gorm.DB
)

//get connection to db, return pointer to struct gorm db
func GetConnectionDb() *gorm.DB {
	once.Do(func() {
		connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.GetInst().DbHost, config.GetInst().DbPort, config.GetInst().DbUsername, config.GetInst().DbName, config.GetInst().DbPassword)
		gormDb, err := gorm.Open("postgres", connString)
		if err != nil {
			log.Fatal("Connect to Db failed: " + err.Error())
		}
		db = gormDb
	})
	return db
}
