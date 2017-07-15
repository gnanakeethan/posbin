package controllers

import (
	"strings"

	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gnanakeethan/posbin/models"
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
	"github.com/gnanakeethan/posbin/src/auth"
)

// ActionController operations for Token
type ActionController struct {
	beego.Controller
}

func (c *ActionController) Prepare() {
	versionString := "/v1/"
	o := orm.NewOrm()
	request := c.Ctx.Request.URL.Path
	request = strings.Replace(request, versionString, "", 1)
	if !strings.HasPrefix(request, "auth") {
		response := responses.Authentication{}
		token := c.Ctx.Request.Header.Get("Authorization")
		if token != "" {
			auth.ValidateToken(requests.AuthenticationRefreshRequest{Token: token}, &response)
		}
		if response.Success {
			claims := auth.ParseToken(token)
			user := models.Users{Id: claims.UserId}
			method := c.Ctx.Request.Method
			var re = regexp.MustCompile(`\/(\d+)`)
			path := c.Ctx.Request.URL.Path
			permissionString := getPermissionString(re, path, method)
			if claims.UserId != 0 && !user.HasRole("superadmin") {
				var permissions []*models.Permissions
				customquery := "select p.* from users u inner join role_user ru on ru.user_id=u.id " +
					"inner join roles r on ru.role_id=r.id " +
					"inner join permission_role rp on rp.role_id=r.id " +
					"inner join permissions p on p.id=rp.permission_id " +
					"where p.name= ? "
				response.Success = false
				if count, err := o.Raw(customquery, permissionString).QueryRows(&permissions); err == nil && count > 0 {
					response.Success = true
				}
			}
		}
		if !response.Success {
			c.Abort("401")
		}

	}
}
func getPermissionString(re *regexp.Regexp, path, method string) string {
	permissionString := strings.ToLower(strings.Replace(strings.TrimPrefix(re.ReplaceAllString(path, ``), "/v1/"), "/", "_", -1) + "_" + method)
	return permissionString
}

func (c *ActionController) GetUser() {
	token := c.Ctx.Request.Header.Get("Authorization")
	if token != "" {
		auth.ParseToken(token)
		//auth.ValidateToken(requests.AuthenticationRefreshRequest{Token: token}, &response)

	}

}
