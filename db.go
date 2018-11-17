package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/twinj/uuid"
)

var db *gorm.DB

func Connect() error {
	addr := viper.GetString("DB_URL")
	fmt.Printf("Connecting to %v\n", addr)

	var err error
	db, err = gorm.Open("postgres", addr)
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&AccountImage{})
	return nil
}

func FindAccountImage(id string) (AccountImage, error) {
	if db == nil {
		return AccountImage{}, fmt.Errorf("DB not initialized")
	}
	tx := db.Begin()
	accountImage := &AccountImage{}
	tx = tx.Find(&accountImage, "ID=?", id)
	if tx.Error != nil {
		return *accountImage, tx.Error
	}
	tx = tx.Commit()
	return *accountImage, nil
}

func ListAccountImages() ([]AccountImage, error) {
	if db == nil {
		return []AccountImage{}, fmt.Errorf("DB not initialized")
	}
	accountImages := make([]AccountImage, 0)
	tx := db.Begin()
	tx.Find(&accountImages)
	if tx.Error != nil {
		return accountImages, tx.Error
	}
	tx = tx.Commit()
	return accountImages, nil
}

func SeedRandomAccountImage() error {
	guid := uuid.NewV4().String()
	tx := db.Begin()
	tx = tx.Create(&AccountImage{ID: guid, URL: "http://callistaenterprise.se", ServedBy: "localhost"})
	if tx.Error != nil {
		return tx.Error
	}
	tx = tx.Commit()
	return nil
}

func health() bool {
	err := db.DB().Ping()
	return err == nil
}
