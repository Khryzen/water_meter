package models

import "github.com/uadmin/uadmin"

type Client struct {
	uadmin.Model
	FirstName     string
	LastName      string
	Address       string
	Landmark      string
	EmailAddress  string
	ContactNumber string
	Active        bool
}
