package utils

import (
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
	var attempt int
	for {
		if db, err := gorm.Open(mysql.Open(configs.DSN), &gorm.Config{}); err != nil {
			attempt++
			if attempt > max_conn {
				return nil, err
			} else {
				log.Printf("DB connection failed, tried %v.", attempt)
				log.Print("Sleeping 5 seconds...")
				time.Sleep(5 * time.Second)
				continue
			}
		} else {
			return db, nil
		}
	}
}

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
