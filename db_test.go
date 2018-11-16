// +build integration

package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/twinj/uuid"
	"net/http"
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
		guid := createTestAccount(t)

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

	deleteTestAccount()
}

func TestApi(t *testing.T) {
	err := Connect()
	defer db.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	SetupGin()

	Convey("Given", t, func() {
		guid := createTestAccount(t)

		Convey("When", func() {
			resp, err := http.Get("http://localhost:8080/accounts/" + guid)
			So(err, ShouldBeNil)
			Convey("Then", func() {
				So(resp.StatusCode, ShouldEqual, 200)
				body, _ := ioutil.ReadAll(resp.Body)
				accountImage := &AccountImage{}
				json.Unmarshal(body, &accountImage)
				So(accountImage.URL, ShouldEqual, SAMPLE_URL)
			})
		})
	})

	deleteTestAccount()
}

func deleteTestAccount() {
	tx := db.Begin()
	tx.Delete(&AccountImage{})
	tx.Commit()
	fmt.Println("Cleaned up!")
}

func createTestAccount(t *testing.T) string {
	guid := uuid.NewV4().String()
	tx := db.Begin()
	tx = tx.Create(&AccountImage{ID: guid, URL: SAMPLE_URL, ServedBy: "localhost"})
	if tx.Error != nil {
		t.Error(fmt.Printf("Error creating AccountImage: %v", tx.Error.Error()))
	}
	tx = tx.Commit()
	return guid
}
