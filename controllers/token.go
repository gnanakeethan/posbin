package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
	"github.com/gnanakeethan/posbin/src/auth"
)

// TokenController operations for Token
type TokenController struct {
	beego.Controller
}

func (c *TokenController) Prepare() {
	versionString := "/v1/"
	request := c.Ctx.Request.RequestURI
	request = strings.Replace(request, versionString, "", 1)
	if !strings.HasPrefix(request, "auth") {
		response := responses.Authentication{}
		token := c.Ctx.Request.Header.Get("Authorization")
		if token != "" {
			auth.ValidateToken(requests.AuthenticationRefreshRequest{Token: token}, &response)

		}

		if response.Success {
		} else {
			c.Abort("401")
		}

	}
}
