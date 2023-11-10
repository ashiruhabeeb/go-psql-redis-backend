package entity

import (
	"time"

	"github.com/google/uuid"
)

// Described user credentials
type User struct {
	UserId		uuid.UUID	`json:"userid"`
	Firstname	string		`json:"firstname" validate:"required,min=2,max=30"`
	Lastname	string		`json:"lastname" validate:"required,min=3,max=30"`
	Username	string		`json:"username" validate:"required,min=2"`
	Email		string		`json:"email" validate:"required,email"`
	Password	string		`json:"password" validate:"required,min=7,containsany=?@!*#"`
	Phone		string		`json:"phone" validate:"required,e164"`
	Createdat	time.Time	`json:"craeted_at"`
	Updatedat	time.Time	`json:"updated_at"`
}

// Describe user address credentials
type Address struct {
	AddressId	uuid.UUID	`json:"addressid"`
	HouseNumber	int			`json:"house_number" validate:"required"`
	StreetName	string		`json:"street_name" validate:"required"`
	LocalArea	string		`json:"lga" validate:"required"`
	State		string		`json:"state" validate:"required"`
	Country		string		`json:"country" validate:"required"`
}

// Beforesave method func auto-generates uuid for user record before user record creation
func (u *User) BeforeUserSave() error {
	u.UserId = uuid.New()

	return nil
}

// Beforesave method func auto-generates uuid for address record before address record creation
func (a *Address) BeforeAddressSave() error {
	a.AddressId = uuid.New()

	return nil
}
