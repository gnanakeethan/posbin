package responses

//Authentication defines the success request of a authenticationrequest
type Authentication struct {
	Success              bool   `json:"success"`
	AuthenticationHeader string `json:"authentication_header"`
}
