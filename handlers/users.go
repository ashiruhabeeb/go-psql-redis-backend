package handlers

import (
	"time"

	"github.com/ashiruhabeeb/go-backend/app/entity"
	"github.com/ashiruhabeeb/go-backend/pkg/password"
	"github.com/ashiruhabeeb/go-backend/pkg/response"
	"github.com/ashiruhabeeb/go-backend/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UsersHandler struct holds users storage repository interface
type UsersHandler struct {
	Storage *storage.UserRepo
}

// NewUsersHandlers creates a new instance of UsersHandlers
func NewUsersHandlers(s *storage.UserRepo) *UsersHandler {
	return &UsersHandler{Storage: s}
}

// Creates a new user based on payload provided by the client side
func(uh *UsersHandler) UserSignUp(c *gin.Context){
	var payload struct {
		Firstname	string			`json:"firstname" validate:"required,min=2,max=30"`
		Lastname	string			`json:"lastname" validate:"required,min=3,max=30"`
		Username	string			`json:"username" validate:"required,min=2"`
		Email		string			`json:"email" validate:"required,email"`
		Password	string			`json:"password" validate:"required,min=7,containsany=?@!*#"`
		ConfirmPassword	string		`json:"confirmpassword" validate:"required"`
		Phone		string			`json:"phone" validate:"required,e164"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 400, err.Error())
		c.Abort()
		return
	}

	if payload.ConfirmPassword != payload.Password {
		response.Error(c, 400, "Password mismatch. Try again!")
		c.Abort()
		return
	}

	hash, err := password.HashPassword(payload.ConfirmPassword)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	// generate a uuid string
	uuidstring := uuid.NewString()

	now := time.Now()
	user := entity.User{
			UserId:	   uuidstring,
			Firstname: payload.Firstname,
			Lastname:  payload.Lastname,
			Username:  payload.Username,
			Email:     payload.Email,
			Password:  hash,
			Phone:     payload.Phone,
			Createdat: now,
	}

	id, err := uh.Storage.InsertUser(user)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.SignupSuccess(c, 201, id)
}

// GetUserById fetch a user's record based on the id parameter provided
func(u *UsersHandler) GetUserById(c *gin.Context){
	id := c.Param("id")

	user, err := u.Storage.FetchUserById(id)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "user record retrieved", user)
}

// GetUserById fetch a user's record based on the id parameter provided
func(u *UsersHandler) GetUserByEmail(c *gin.Context){
	email := c.Query("email")

	user, err := u.Storage.FetchUserByEmail(email)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "user record retrieved", user)
}

// GetUserById fetch a user's record based on the id parameter provided
func(u *UsersHandler) GetUserByUsername(c *gin.Context){
	username := c.Query("username")

	user, err := u.Storage.FetchUserByEmail(username)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "user record retrieved", user)
}

// FetchAllUsersRecords  retrieves all users records in the users table
func(u *UsersHandler) FetchAllUsersRecords(c *gin.Context){
	users, err := u.Storage.FetchAllUsers()
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "Fetched all records", users)
}
