package entity

import (
	"time"

	"github.com/google/uuid"
)

// Described user credentials
type User struct {
	UserId		string		`json:"userid" validate:"required"`
	Firstname	string		`json:"firstname" validate:"required,min=2,max=30"`
	Lastname	string		`json:"lastname" validate:"required,min=3,max=30"`
	Username	string		`json:"username" validate:"required,min=2"`
	Email		string		`json:"email" validate:"required,email"`
	Password	string		`json:"password" validate:"required,min=7,containsany=?@!*#"`
	Phone		string		`json:"phone" validate:"required,e164"`
	Address		Address		`json:"address" validate:"dive"`
	CreatedOn	time.Time	`json:"craetedon"`
	UpdatedOn	time.Time	`json:"updatedon"`
}

// Describe user address credentials
type Address struct {
	HouseNumber	int		`json:"house_number" validate:"required"`
	StreetName	string	`json:"street_number" validate:"required"`
	LocalArea	string	`json:"lga" validate:"required"`
	State		string	`json:"state" validate:"required"`
	Country		string	`json:"country" validate:"required"`
}

// Beforesave method func auto-generates uuid string for user record before user record creation
func (u *User) Beforesave() error {
	u.UserId = uuid.NewString()

	return nil
}
