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

//Authenticate function defines authentication for user;
func Authenticate(v requests.AuthenticationRequest, response *responses.Authentication) {
	o := orm.NewOrm()
	o.QueryTable(new(models.Users)).Filter("username", v.Username).One(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password)); err == nil {
		// Create the Claims
		claims := AuthenticationClaim{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 3600,
				NotBefore: time.Now().Unix(),
			},
		}
		signingString := user.Username + user.Email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if response.Token, err = token.SignedString([]byte(signingString)); err == nil {
			response.Success = true
		}
	}
	return
}
func ValidateToken(v requests.AuthenticationRefreshRequest, response *responses.Authentication) {
	logs.Info(time.Local)
	response.Token = ""
	claims := AuthenticationClaim{}
	token, _ := jwt.ParseWithClaims(v.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		o := orm.NewOrm()
		vp := &models.Users{Id: claims.UserId}
		if err := o.Read(vp); err == nil {
			return []byte(vp.Username + vp.Email), nil
		}
		return []byte(""), nil
	})

	response.Success = false
	if _, ok := token.Claims.(*AuthenticationClaim); ok && token.Valid && extendedValidation(v.Token) && permissionCheck(claims.UserId) {
		response.Token = v.Token
		response.Success = true
	}
}

//permissionCheck to be implemneted
func permissionCheck(userID int) bool {

	return true
}
func extendedValidation(token string) bool {
	o := orm.NewOrm()
	if count, err := o.QueryTable(new(models.InvalidTokens)).Filter("token", token).Count(); err == nil {
		if count > 0 {
			return false
		}
	}
	return true
}
func RefreshToken(v requests.AuthenticationRefreshRequest, response *responses.Authentication) {
	response.Success = false
	response.Token = ""
	claims := AuthenticationClaim{}
	token, _ := jwt.ParseWithClaims(v.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		o := orm.NewOrm()
		vp := &models.Users{Id: claims.UserId}
		if err := o.Read(vp); err == nil {
			return []byte(vp.Username + vp.Email), nil
		}
		return []byte(""), nil
	})
	if claims, ok := token.Claims.(*AuthenticationClaim); ok && token.Valid && extendedValidation(v.Token) {
		o := orm.NewOrm()
		go func(v requests.AuthenticationRefreshRequest) {
			_, offset := time.Now().Zone()
			destroyToken := models.InvalidTokens{Token: v.Token, ValidThru: time.Unix(claims.ExpiresAt+int64(offset), 200)}
			o.Insert(&destroyToken)
		}(v)

		vp := &models.Users{Id: claims.UserId}
		if err := o.Read(vp); err == nil {
			claims := AuthenticationClaim{
				UserId: vp.Id,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Unix() + 3600,
					NotBefore: time.Now().Unix(),
				},
			}
			signingString := vp.Username + vp.Email
			token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			if response.Token, err = token.SignedString([]byte(signingString)); err == nil {
				response.Success = true
			}
		}

	}
}
