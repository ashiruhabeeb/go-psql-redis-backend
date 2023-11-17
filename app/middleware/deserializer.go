package middleware

// import (
// 	"context"

// 	"github.com/ashiruhabeeb/go-backend/app/entity"
// 	"github.com/ashiruhabeeb/go-backend/db"
// 	"github.com/ashiruhabeeb/go-backend/pkg/config"
// 	"github.com/ashiruhabeeb/go-backend/pkg/response"
// 	"github.com/ashiruhabeeb/go-backend/pkg/token"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"github.com/redis/go-redis/v9"
// )

// func DeserializeUser(c *gin.Context) {
// 	tokenString := c.GetHeader("Authorization")

// 	if tokenString == "" {
// 		response.Error(c, 401, "authorization header is required")
// 	}

// 	tkn := tokenString[len("Bearer "):]
// 	if tkn == "" {
// 		response.Error(c, 401, "authorization header is required")
// 	}

// 	tknClaims, err := token.ValidateToken(tkn, config.Cfg.AccessTokenPublicKey)
// 	if err != nil {
// 		response.Error(c, 403, err.Error())
// 		c.Abort()
// 	}

// 	ctx := context.TODO()
// 	userId, err := db.RedisClient.Get(ctx, tknClaims["sub"]).Result()
// 	if err == redis.Nil {
// 		response.Error(c, 403, err.Error())
// 		c.Abort()
// 	}

// 	uuid, err := uuid.Parse(userId)
// 	if err != nil {
// 		response.Error(c, 500, err.Error())
// 		c.Abort()
// 	}

// 	var e = entity.User{}

// 	row := db.Db.QueryRow(db.PsqlFetchUserById, uuid)

// 	err = row.Scan(&e.UserId, &e.Firstname, &e.Lastname, &e.Username, &e.Email, &e.Phone, &e.Createdat, &e.Updatedat)
// 	if err != nil {
// 		if err == row.Err(){
// 			return 
// 		}
// 		return
// 	}

// 	c.Set("user", e)
// 	c.Set("access_token_uuid", tknClaims)
// 	c.Next()
// }