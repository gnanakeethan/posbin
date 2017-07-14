package controllers

import (
	"strings"

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
			paths := strings.Split(c.Ctx.Request.URL.Path, "/")
			permissionString := strings.ToLower(paths[0] + "_" + method)
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

func (c *ActionController) GetUser() {
	token := c.Ctx.Request.Header.Get("Authorization")
	if token != "" {
		auth.ParseToken(token)
		//auth.ValidateToken(requests.AuthenticationRefreshRequest{Token: token}, &response)

	}

}
