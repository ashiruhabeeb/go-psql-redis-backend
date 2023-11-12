package token

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CreateToken(userId, secretkey string, ttl time.Duration)(*TokenClaims, error){
	var td = &TokenClaims{
		ExpiresIn: new(int64),
		Token: new(string),
	}

	*td.ExpiresIn = time.Now().UTC().Add(ttl).Unix()
	td.TokenUUID = uuid.NewString()
	td.UserID = userId

	decodedSecretKey, err := base64.StdEncoding.DecodeString(secretkey)
	if err != nil {
		return nil, fmt.Errorf("token secret key decode failure: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedSecretKey)
	if err != nil {
		return nil, fmt.Errorf("parse decodedSecretKey failure: %w", err)
	}

	atClaims := make(jwt.MapClaims)
	atClaims["sub"] = userId
	atClaims["token_uuid"] = td.TokenUUID
	atClaims["exp"] = td.ExpiresIn
	atClaims["iat"] = time.Now().UTC().Unix()
	atClaims["nbf"] = time.Now().UTC().Unix()

	*td.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims).SignedString(privKey)
	if err != nil {
		return nil, fmt.Errorf("create: sign token: %w", err)
	}

	return td, nil
}