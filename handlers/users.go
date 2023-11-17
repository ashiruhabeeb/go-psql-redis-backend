package handlers

import (
	"github.com/ashiruhabeeb/go-backend/app/entity"
	"github.com/ashiruhabeeb/go-backend/pkg/config"
	"github.com/ashiruhabeeb/go-backend/pkg/response"
	"github.com/ashiruhabeeb/go-backend/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UsersHandler struct holds users storage repository interface
type UsersHandler struct {
	Storage *storage.UserRepo
	cfg	*config.Config
}

// NewUsersHandlers creates a new instance of UsersHandlers
func NewUsersHandlers(s *storage.UserRepo, cfg	*config.Config) *UsersHandler {
	return &UsersHandler{Storage: s, cfg: cfg}
}

// GetUserById fetch a user's record based on the id parameter provided
func(u *UsersHandler) GetUserById(c *gin.Context){
	userid := c.Param("userid")

	uuid, err := uuid.Parse(userid)
	HandleError(err)

	user, err := u.Storage.FetchUserById(uuid)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "user record retrieved", user)
}

// GetUserById fetch a user's record based on the id parameter provided
func(u *UsersHandler) GetUserByEmail(c *gin.Context){
	email := c.Param("email")

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
	username := c.Param("username")

	user, err := u.Storage.FetchUserByUsername(username)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "user record retrieved", user)
}

// FetchAllUsersRecords retrieves all users records in the users table
func(u *UsersHandler) FetchAllUsersRecords(c *gin.Context){
	users, err := u.Storage.FetchAllUsers()
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "Fetched all records", users)
}

// UpdateUser updates user's record fron the user's table
func(u *UsersHandler) UpdateUser(c *gin.Context){
	var payload struct {
		Firstname		string	`json:"firstname" validate:"min=2,max=30"`
		Lastname		string	`json:"lastname" validate:"min=3,max=30"`
		Username		string	`json:"username" validate:"min=2"`
		Phone			string	`json:"phone" validate:"required,e164"`
	}
	
	userid := c.Param("userid")

	uuid , err := uuid.Parse(userid)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 400, err.Error())
		c.Abort()
		return
	}

	if err = u.Storage.UpdateUserRecord(uuid, payload.Firstname, payload.Lastname, payload.Username, payload.Phone); err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "user record updated successfully", nil)
}

// DeleteUser deletes user's record from users table based 
func(u *UsersHandler) DeleteUser(c *gin.Context){
	userid := c.Param("userid")

	uuid, err := uuid.Parse(userid)
	HandleError(err)

	e := entity.User{UserId: uuid}

	err = u.Storage.DeleteUser(e.UserId)
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	response.Success(c, 200, "User successfully deleted", nil)
}
