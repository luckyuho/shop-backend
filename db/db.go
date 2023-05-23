package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		logrus.Error(err)
	}
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST")+":"+os.Getenv("DBPORT"),
		os.Getenv("DBNAME")))
	fmt.Println(os.Getenv("DBUSER"), ":", os.Getenv("DBPASSWORD"), "@tcp(", os.Getenv("DBHOST")+":"+os.Getenv("DBPORT"), ")/", os.Getenv("DBNAME"))

	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func Get() *gorm.DB {
	return db
}
