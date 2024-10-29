package utils

import (
	"fmt"
	"local/fin/configs"
	"local/fin/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connDB(max_conn int) (*gorm.DB, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered: %v", r)
		}
	}()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", configs.DB_USER, configs.DB_PASS, configs.DB_HOST, configs.DB_PORT)
	log.Print(dsn)
	for i := 0; max_conn > i; i++ {
		if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			log.Printf("DB connection failed, tried %v.", i+1)
			log.Print("Sleeping 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		} else {
			_ = db.Exec("CREATE DATABASE IF NOT EXISTS myapp;")
		}
	}
	
	db, _ := gorm.Open(mysql.Open(dsn + configs.DB_TABLE), &gorm.Config{})
	return db, nil
}
// 	for {
// 		if err != nil {
// 			panic("cannot create table")
// 		} else {
// 			_ = database.Exec("CREATE DATABASE IF NOT EXISTS myapp;")
// 		}
// 		if db, err := gorm.Open(mysql.Open(dsn + configs.DB_TABLE), &gorm.Config{}); err != nil {
// 			attempt++
// 			if attempt > max_conn {
// 				return nil, err
// 			} else {
// 				log.Printf("DB connection failed, tried %v.", attempt)
// 				log.Print("Sleeping 5 seconds...")
// 				time.Sleep(5 * time.Second)
// 				continue
// 			}
// 		} else {
// 			return db, nil
// 		}
// 	}
// }

func GetDB() *gorm.DB {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered: %v", r)
		}
	}()
	db, err := connDB(configs.CONN_TRY_COUNT)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic(err)
	}
	log.Print("Successfuly connect to db")
	return db
}
