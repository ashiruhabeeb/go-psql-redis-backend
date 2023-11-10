package storage

import (
	"database/sql"
	"fmt"

	"github.com/ashiruhabeeb/go-backend/app/entity"
	"github.com/ashiruhabeeb/go-backend/db"
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
func (u *UserRepo) InsertUser(e entity.User) (string, error) {
	err := u.db.QueryRow(db.PsqlInsertUser, e.UserId, e.Firstname, e.Lastname, e.Username, e.Email, 
		e.Password, e.Phone).Scan(&e.UserId)

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return e.UserId, nil
}
	

// FetchUserById retrieves a specific user record from the users table
func (u *UserRepo) FetchUserById(userid string) (*entity.User, error){
	e := entity.User{}

	row := u.db.QueryRow(db.PsqlFetchUserByEmail, userid)

	err := row.Scan(&e.Firstname, &e.Lastname, &e.Username, &e.Email, &e.Phone, &e.Createdat, &e.Updatedat)
	if err != nil {
		if err == row.Err() {
			return nil, err
		}
		return nil, err
	}

	return &e, nil
}

// FetchUserById retrieves a specific user record from the users table based on userid as parameter
func (u *UserRepo) FetchUserByEmail(email string) (*entity.User, error){
	e := entity.User{}

	row := u.db.QueryRow(db.PsqlFetchUserByEmail, email)

	err := row.Scan(&e.UserId, &e.Firstname, &e.Lastname, &e.Username, &e.Email, &e.Phone, &e.Createdat, &e.Updatedat)
	if err != nil {
		if err == row.Err() {
			return nil, err
		}
		return nil, err
	}

	return &e, nil
}

// FetchUserByUsername retrieves a specific user record from the users table based on userid as parameter
func (u *UserRepo) FetchUserByUsername(username string) (*entity.User, error){
	e := entity.User{}

	row := u.db.QueryRow(db.PsqlFetchUserByEmail, username)

	err := row.Scan(&e.UserId, &e.Firstname, &e.Lastname, &e.Username, &e.Email, &e.Phone, &e.Createdat, &e.Updatedat)
	if err != nil {
		if err == row.Err() {
			return nil, err
		}
		return nil, err
	}

	return &e, nil
}

// FetchAllUsers retrieves all user records present in users table in groups of 3
func(u *UserRepo) FetchAllUsers()([]entity.User, error){
	rows, err := u.db.Query(db.PsqlFetchUsers)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer rows.Close()

	users := []entity.User{}
	for rows.Next(){
		var user entity.User
		err := rows.Scan(&user.UserId, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.Phone, &user.Createdat, &user.Updatedat)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

// Implements UserStorage methods 
type UserRepository interface{
	InsertUser(e entity.User) (string, error)
	FetchUserById(userid string) (*entity.User, error)
	FetchUserByEmail(email string) (*entity.User, error)
	FetchUserByUsername(username string) (*entity.User, error)
	FetchAllUsers()([]entity.User, error)
}
