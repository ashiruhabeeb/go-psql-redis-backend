package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Described user credentials
type User struct {
	UserId		string	
	Firstname	string		
	Lastname	string		
	Username	string		
	Email		string		
	Password	string		
	Phone		string		
	Createdat	time.Time	
	Updatedat	*sql.NullTime
}

// Describe user address credentials
type Address struct {
	AddressId	string	`json:"addressid"`
	HouseNumber	int		`json:"house_number" validate:"required"`
	StreetName	string	`json:"street_name" validate:"required"`
	LocalArea	string	`json:"lga" validate:"required"`
	State		string	`json:"state" validate:"required"`
	Country		string	`json:"country" validate:"required"`
}

// Beforesave method func auto-generates uuid for user record before user record creation
func (u *User) BeforeUserSave() error {
	u.UserId = uuid.NewString()

	return nil
}

// Beforesave method func auto-generates uuid for address record before address record creation
func (a *Address) BeforeAddressSave() error {
	a.AddressId = uuid.NewString()

	return nil
}
