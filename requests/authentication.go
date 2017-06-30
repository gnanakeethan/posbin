package requests

//AuthenticationRequest is used to receive post values for authentication
type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticationRefreshRequest struct {
	Token string `json:"token"`
}

//Validate function defined the validation functionality for request
func (request AuthenticationRequest) Validate() bool {
	if (request.Username != "") && (request.Password != "") {
		return true
	}
	return false
}

//Validate function defined the validation functionality for request
func (request AuthenticationRefreshRequest) Validate() bool {
	if request.Token != "" {
		return true
	}
	return false
}
