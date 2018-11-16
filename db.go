package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func Connect() *gorm.DB {
	addr := viper.GetString("DB_URL")
	fmt.Printf("Connecting to %v\n", addr)
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&AccountImage{})
	return db
}
