package auth

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gnanakeethan/posbin/models"
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
	"golang.org/x/crypto/bcrypt"
)

var user models.Users

type MyCustomClaims struct {
	UserId int `json:"UserId"`
	jwt.StandardClaims
}

//Authenticate function defines authentication for user;
func Authenticate(v requests.AuthenticationRequest, response *responses.Authentication) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Users))
	qs.Filter("username", v.Username)
	qs.One(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password)); err == nil {
		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		logs.Info(user)

		// Create the Claims
		claims := MyCustomClaims{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: 15000,
				NotBefore: time.Now().Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//What the hell is this?
		mySigningKey := []byte("AllYourBase")
		tokenString, _ := token.SignedString(mySigningKey)
		response.Success = true
		response.AuthenticationHeader = tokenString
	}

}

func validUser(username string, password string) bool {

	return false
}

func getRoles() {

}
