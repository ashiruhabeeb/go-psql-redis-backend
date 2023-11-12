package token

type TokenClaims struct {
	Token		*string
	TokenUUID	string
	UserID		string
	ExpiresIn	*int64
}
