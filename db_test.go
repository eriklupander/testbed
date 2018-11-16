// +build integration

package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/twinj/uuid"
	"testing"
)

var SAMPLE_URL = "http://callistaenterprise.se"

func TestDb(t *testing.T) {
	err := Connect()
	defer db.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	Convey("Given", t, func() {
		// Create something in the DB
		guid := uuid.NewV4().String()
		tx := db.Begin()
		tx = tx.Create(&AccountImage{ID: guid, URL: SAMPLE_URL, ServedBy: "localhost"})
		if tx.Error != nil {
			t.Error(fmt.Printf("Error creating AccountImage: %v", tx.Error.Error()))
		}
		tx = tx.Commit()

		Convey("When", func() {
			// Load it
			acc := &AccountImage{}
			tx := db.Begin()
			tx = tx.First(&acc, "ID = ?", guid)
			tx = tx.Commit()
			Convey("Then", func() {
				// Assert
				So(acc.URL, ShouldEqual, SAMPLE_URL)
			})
		})
	})

	tx := db.Begin()
	tx.Delete(&AccountImage{})
	tx.Commit()
	fmt.Println("Cleaned up!")
}
