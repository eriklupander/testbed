// +build testtools

package main

import (
	"fmt"
	"github.com/twinj/uuid"
	"testing"
)

var SAMPLE_URL = "http://callistaenterprise.se"

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
