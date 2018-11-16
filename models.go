package main

// AccountImage with GORM ID
type AccountImage struct {
	ID       string `json:"id" gorm:"primary_key"`
	URL      string `json:"url"`
	ServedBy string `json:"servedBy"`
}
