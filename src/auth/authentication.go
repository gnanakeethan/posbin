package auth

import (
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
)

//Authenticate function defines authentication for user;
func Authenticate(v requests.AuthenticationRequest, response *responses.Authentication) (key string, success bool) {
	return "", true
}
