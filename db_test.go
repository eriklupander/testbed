// +build integration

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDb(t *testing.T) {
	err := Connect()
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	Convey("Given there is a known row in the DB", t, func() {
		// Create something in the DB
		guid := createTestAccount(t)

		Convey("When read from the db", func() {
			// Load it
			acc := &AccountImage{}
			tx := db.Begin()
			tx = tx.First(&acc, "ID = ?", guid)
			tx = tx.Commit()
			Convey("Then we should get the correct value", func() {
				// Assert
				So(tx.Error, ShouldBeNil)
				So(acc.ID, ShouldEqual, guid)
				So(acc.URL, ShouldEqual, SAMPLE_URL)
				So(acc.ServedBy, ShouldEqual, "localhost")
			})
		})

		Convey("When trying to read a non-existing object", func() {
			// Load it
			acc := &AccountImage{}
			tx := db.Begin()
			tx = tx.First(&acc, "ID = ?", "nonexisting")
			tx = tx.Commit()
			Convey("Then we should get an error", func() {
				// Assert
				So(tx.Error, ShouldNotBeNil)
			})
		})

	})

	deleteTestAccount()
}
