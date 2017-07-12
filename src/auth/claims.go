package auth

import jwt "github.com/dgrijalva/jwt-go"

type AuthenticationClaim struct {
	UserId int `json:"UserId"`

	jwt.StandardClaims
}
