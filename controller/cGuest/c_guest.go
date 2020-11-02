package cGuest

import (
	"be01gofire/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK,`guest_login.html`,gin.H{
			`foo`:`bar`,
		})
		return
	}
	// TODO: ambil inputan (bind ke struct dari form)
	// TODO: panggil model untuk login, misal: mUsers.Login(...)
	res := map[string]interface{}{}
	res[`ok`] = true
	ctx.JSON(http.StatusOK,res)
}
