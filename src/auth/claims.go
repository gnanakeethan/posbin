package auth

import jwt "github.com/dgrijalva/jwt-go"

type AuthenticationClaim struct {
	UserId     int `json:"UserId"`
	StoreId    int `json:"StoreId"`
	TerminalId int `json:"TerminalId"`
	jwt.StandardClaims
}
