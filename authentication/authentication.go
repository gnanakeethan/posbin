package authentication

import (
	"fmt"
)

//AuthenticationRequest is used to receive post values for authentication
type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Validate function defined the validation functionality for request
func (request AuthenticationRequest) Validate() bool {
	if (request.Username != "") && (request.Password != "") {
		return true
	}
	return false
}

//Authenticate function defines authentication for user;
func Authenticate() {
	fmt.Println("test")
}
