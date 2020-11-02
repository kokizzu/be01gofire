package cGuest

import (
	"be01gofire/controller"
	"be01gofire/model/mUser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_login.html`, gin.H{
			`foo`: `bar`,
		})
		return
	}
	user := mUser.User{}
	err := ctx.ShouldBindJSON(&user)
	pass := user.Pass
	res := map[string]interface{}{}
	if err == nil {
		err := user.Find(ctx.Firestore)
		if err == nil {
			if user.CheckPass(pass) {
				res[`email`] = user.Email
				// TODO: set cookie/jwt
			} else {
				res[`error`] = `wrong username or password` // wrong password
			}
		} else {
			res[`error`] = `wrong username or password` // not exists
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func Register(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_register.html`, gin.H{})
		return
	}
	// https://github.com/gin-gonic/gin#bind-form-data-request-with-custom-struct 
	user := mUser.User{}
	err := ctx.ShouldBindJSON(&user)
	res := map[string]interface{}{}
	if err == nil {
		user.HashPassword()
		err = user.Insert(ctx.Firestore)
		if err == nil {
			res[`user`] = user
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}
