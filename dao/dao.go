package dao

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	url := fmt.Sprintf("%s:%s@(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	var err error
	db, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
		log.Println("failed to init dao")
	}
	db.LogMode(true)
	log.Println("success to init dao")
}
