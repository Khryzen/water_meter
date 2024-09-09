package models

import "github.com/uadmin/uadmin"

type Device struct {
	uadmin.Model
	SerialNumber string
	Deployed     bool
	Active       bool
}

func (d *Device) String() string {
	return d.SerialNumber
}
