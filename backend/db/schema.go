package db

import (
	"time"

	"gorm.io/gorm"
)

type Prospect struct {
	gorm.Model
	CustomerID  int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	CompanyName string
	Website     string
	Notes       string
	DateCreated time.Time
	DateUpdated time.Time
}
