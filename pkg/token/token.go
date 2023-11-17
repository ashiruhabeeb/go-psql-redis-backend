package token

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	issuer string
}

func NewjwtService(issuer string) *jwtService{
	return &jwtService{issuer: issuer}
}

type JWTService interface {
	CreateToken(privatekey string, payload interface{}, ttl time.Duration)(string, error)
	ValidateToken(token, publicKey string)(TokenClaims, error)
}

// CreateToken func generates JWT token for user authorization
func CreateToken(privatekey string, payload interface{}, ttl time.Duration)(string, error){
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privatekey)
	if err != nil {
		return "", fmt.Errorf("token secret key decode failure: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("parse decodedPrivateKey failure: %w", err)
	}

	atClaims := make(jwt.MapClaims)
	atClaims["sub"] = payload
	atClaims["exp"] = time.Now().Add(ttl).Unix()
	atClaims["iat"] = time.Now().UTC().Unix()
	atClaims["nbf"] = time.Now().UTC().Unix()

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return accessToken, nil
}

// ValidateToken func validate's user's request
func ValidateToken(token, publicKey string)(interface{}, error){
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("publicKey decode failure: %w", err)
	}

	key, err  := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return "", fmt.Errorf("parse decodedPublicKey failure: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token validation error: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
