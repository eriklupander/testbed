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
