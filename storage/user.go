package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ashiruhabeeb/go-backend/app/entity"
	"github.com/ashiruhabeeb/go-backend/db"
)

// UserStorage holds db object of type database/sql package
type UserStorage struct {
	db *sql.DB
}

// NewUserStorage constructor creates a new instance of UserStorage object
func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db}
}

// InsertUser creates a new user record in the users table
func (u *UserStorage) InsertUser(e entity.User)(string, error){
	err := u.db.QueryRow(db.PsqlInsertUser, e.Firstname, e.Lastname, e.Username, e.Email, e.Password, e.Phone, e.Address.HouseNumber, e.Address.StreetName, e.Address.LocalArea, e.Address.State, e.Address.Country).Scan(&e.UserId)
	if err != nil {
		return "", err
	}
	
	return e.UserId, nil
}

// FetchUserById retrieves a specific user record from the users table
func (u *UserStorage) FetchUserById(userid string) (*entity.User, error){
	e := entity.User{}

	row := u.db.QueryRow(db.PsqlFetchUserById, e.UserId)
	switch err := row.Scan(&e); err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
	case nil:
		fmt.Println(userid)
	default:
		log.Println(err)
	}
	return &e, nil
}

// Implements UserStorage methods 
type UserRepo interface{
	InsertUser(payload entity.User)(string, error)
	FetchUserById(userid string) (*entity.User, error)
}
