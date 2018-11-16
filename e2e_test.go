// +build e2e

package main

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"testing"
)

// E2E test, requires running container.
func TestRunningApi(t *testing.T) {
	err := Connect()
	defer db.Close()
	if err != nil {
		t.Errorf(err.Error())
		panic(err.Error())
	}

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
