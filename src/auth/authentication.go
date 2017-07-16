package auth

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gnanakeethan/posbin/models"
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
	"golang.org/x/crypto/bcrypt"
)

var user models.Users
var validTime int64 = 600

//Authenticate function defines authentication for user;
func Authenticate(v requests.AuthenticationRequest, response *responses.Authentication) {
	o := orm.NewOrm()
	o.QueryTable(new(models.Users)).Filter("username", v.Username).One(&user)
	o.LoadRelated(&user, "TerminalId")
	//logs.Info(reflect.DeepEqual(user.TerminalId) == reflect.TypeOf(&models.Terminals{}))
	//panic(user.TerminalId)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password)); err == nil {
		// Create the Claims
		claims := AuthenticationClaim{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + validTime,
				NotBefore: time.Now().Unix(),
			},
		}
		logs.Info(claims)

		signingString := user.Username + user.Email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if response.Token, err = token.SignedString([]byte(signingString)); err == nil {
			response.Success = true
		}

		go func() {
			destroyToken := models.Tokens{Token: response.Token, ValidThru: time.Unix(claims.ExpiresAt, 200), Valid: true, User: &user}
			o.ReadOrCreate(&destroyToken, "valid")

			sqlnull := "update terminals set user_id null where user_id=?"
			o.Raw(sqlnull, user.Id).Exec()
		}()
	}
	return
}

func ParseToken(token string) *AuthenticationClaim {
	claims := AuthenticationClaim{}
	jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		o := orm.NewOrm()
		vp := &models.Users{Id: claims.UserId}
		if err := o.Read(vp); err == nil {
			return []byte(vp.Username + vp.Email), nil
		}
		return []byte(""), nil
	})
	//if _, ok := tokenDecoded.Claims.(*AuthenticationClaim); ok && tokenDecoded.Valid && extendedValidation(token) {
	//	return claims
	//}
	return &claims
}
func ClearTokens() {
	o := orm.NewOrm()

	sql := "delete from tokens where valid_thru < now()"
	_, err := o.Raw(sql).Exec()
	logs.Error(err)

}

func InvalidateToken(v requests.AuthenticationRefreshRequest) {
	o := orm.NewOrm()
	query := "update tokens set `valid` = 0 where `key` = ?;"
	o.Raw(query, v.Token).Exec()
	logs.Info(v.Token)
}

func ValidateToken(v requests.AuthenticationRefreshRequest, response *responses.Authentication) {
	logs.Info(time.Local)
	response.Token = ""
	response.Success = false
	claims := AuthenticationClaim{}
	token, _ := jwt.ParseWithClaims(v.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		o := orm.NewOrm()
		vp := &models.Users{Id: claims.UserId}
		if err := o.Read(vp); err == nil {
			return []byte(vp.Username + vp.Email), nil
		}
		return []byte(""), nil
	})

	if _, ok := token.Claims.(*AuthenticationClaim); ok && token.Valid && extendedValidation(v.Token) {
		response.Token = v.Token
		response.Success = true
	}
}

func extendedValidation(token string) bool {
	o := orm.NewOrm()
	if count, err := o.QueryTable(new(models.Tokens)).Filter("token", token).Filter("valid", false).Count(); err == nil {
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
		vp := &models.Users{Id: claims.UserId}
		go InvalidateToken(v)

		if err := o.Read(vp); err == nil {
			claims := AuthenticationClaim{
				UserId: vp.Id,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Unix() + validTime,
					NotBefore: time.Now().Unix(),
				},
			}
			signingString := vp.Username + vp.Email
			token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			if response.Token, err = token.SignedString([]byte(signingString)); err == nil {
				response.Success = true
			}
			go func() {
				destroyToken := models.Tokens{Token: response.Token, ValidThru: time.Unix(claims.ExpiresAt, 200), Valid: true, User: vp}
				o.Insert(&destroyToken)
			}()
		}

	}
}
