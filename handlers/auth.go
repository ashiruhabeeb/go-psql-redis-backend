package handlers

import (
	"strings"
	"time"

	"github.com/ashiruhabeeb/go-backend/app/entity"
	"github.com/ashiruhabeeb/go-backend/pkg/password"
	"github.com/ashiruhabeeb/go-backend/pkg/response"
	"github.com/ashiruhabeeb/go-backend/pkg/token"
	"github.com/ashiruhabeeb/go-backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Creates a new user based on payload provided by the client side
func(uh *UsersHandler) UserSignUp(c *gin.Context){
	var payload struct {
		Firstname		string	`json:"firstname" validate:"required,min=2,max=30"`
		Lastname		string	`json:"lastname" validate:"required,min=3,max=30"`
		Username		string	`json:"username" validate:"required,min=2"`
		Email			string	`json:"email" validate:"required,email"`
		Password		string	`json:"password" validate:"required,min=7,containsany=?@!*#"`
		ConfirmPassword	string	`json:"confirmpassword" validate:"required"`
		Phone			string	`json:"phone" validate:"required,e164"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 400, err.Error())
		c.Abort()
		return
	}

	if err := validator.Validate(payload); err != nil {
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
	uuidstring := uuid.New()

	now := time.Now()
	user := entity.User{
			UserId:	   uuidstring,
			Firstname: payload.Firstname,
			Lastname:  payload.Lastname,
			Username:  payload.Username,
			Email:     strings.ToLower(payload.Email),
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

	accessTokenDetails, err := token.CreateToken(uh.cfg.AccessTokenPrivateKey, user.UserId.String(), time.Duration(uh.cfg.AccessTokenExpiresIn * int(time.Minute)))
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	refreshAccessTokenDetails, err := token.CreateToken(uh.cfg.RefreshTokenPrivateKey, user.UserId.String(), time.Duration(uh.cfg.RefreshTokenExpiresIn * int(time.Hour)))
	if err != nil {
		response.Error(c, 500, err.Error())
		c.Abort()
		return
	}

	// accessErr := db.RedisClient.Set(context.TODO(), accessTokenDetails.TokenUUID, user.UserId.String(), time.Duration(accessTokenDetails.ExpiresIn)).Err()
	// if accessErr != nil {
	// 	response.Error(c, 422, accessErr.Error())
	// 	c.Abort()
	// 	return
	// }

	// resfreshAccessError := db.RedisClient.Set(context.TODO(), refreshAccessTokenDetails.TokenUUID, user.UserId.String(), time.Duration(refreshAccessTokenDetails.ExpiresIn)).Err()
	// if resfreshAccessError != nil {
	// 	response.Error(c, 422, err.Error())
	// 	c.Abort()
	// 	return
	// }

	c.SetCookie("Authorization", accessTokenDetails, uh.cfg.AccessTokenMaxAge, "/", "localhost", false, true)
	c.SetCookie("Authorization", refreshAccessTokenDetails, uh.cfg.RefreshTokenMaxAge, "/", "localhost", false, true)

	response.Success(c, 201, "user record inserted", id)
}

// Sign in an exxisting user based on the provided payload
func(uh *UsersHandler) SignIn(c *gin.Context) {

} 
