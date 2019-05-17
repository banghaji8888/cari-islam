package db

import (
	"cari-islam/config"
	"cari-islam/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // driver
)

var db *gorm.DB
var err error

// Init database connection
func Init() {
	conf := config.GetConfig()
	var host string
	if conf.IsXampp {
		host = conf.DBHostXampp
	}

	connectString := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", conf.DBUsername, conf.DBPassword, host, conf.DBName)

	db, err = gorm.Open("mysql", connectString)

	if err != nil {
		log.Println(err)
		panic("DB Connection Error")
	}
	db.SingularTable(true)
	db.AutoMigrate(&model.Sumber{}, &model.Kategori{}, &model.Konten{})

}

// GetConnection get current database connection
func GetConnection() *gorm.DB {
	return db
}
