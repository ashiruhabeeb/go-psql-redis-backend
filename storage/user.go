package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ashiruhabeeb/go-backend/app/entity"
	"github.com/ashiruhabeeb/go-backend/db"
	"github.com/google/uuid"
)

// UserStorage holds db object of type database/sql package
type UserRepo struct {
	db *sql.DB
}

// NewUserStorage constructor creates a new instance of UserStorage object
func NewUserStorage(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// InsertUser creates a new user record in the users table
func (u *UserRepo) InsertUser(e entity.User) (uuid.UUID, error) {
	err := u.db.QueryRow(db.PsqlInsertUser, e.Firstname, e.Lastname, e.Username, e.Email, 
		e.Password, e.Phone, e.Address.HouseNumber, e.Address.StreetName, e.Address.LocalArea, 
		e.Address.State, e.Address.Country).Scan(e.UserId)

	if err != nil {
		return uuid.Nil, err
	}
	return e.UserId, nil
}
	

// FetchUserById retrieves a specific user record from the users table
func (u *UserRepo) FetchUserById(userid string) (*entity.User, error){
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
type UserRepository interface{
	InsertUser(e entity.User) (uuid.UUID, error)
	FetchUserById(userid string) (*entity.User, error)
}
