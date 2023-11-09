package handlers

import (
	"github.com/ashiruhabeeb/go-backend/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UsersHandler struct holds users storage repository interface
type UsersHandler struct {
	repo *storage.UserRepo
}

// NewUsersHandlers creates a new instance of UsersHandlers
func NewUsersHandlers(repo storage.UserRepo) *UsersHandler {
	return &UsersHandler{repo: &repo}
}

// Creates a new user based on payload provided by the client side
func(u *UsersHandler) UserSignUP(c *gin.Context){
	id := uuid.NewString()
	c.JSON(201, gin.H{"data": id})
}

// Retrieves a single record of a user
func(u *UsersHandler) GetUserById(c *gin.Context){
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
		return
	}
	c.JSON(200, gin.H{"status": id})
}
