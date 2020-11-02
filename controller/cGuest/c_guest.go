package cGuest

import (
	"be01gofire/controller"
	"be01gofire/model/mUser"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
		err := user.FindByEmail(ctx.Firestore)
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
		ctx.HTML(http.StatusOK, `guest_register.html`, gin.H{
			`title`:`Register ` + time.Now().Format(time.RFC1123Z), 
		})
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

func AddQueue(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_add-queue.html`, gin.H{
			`foo`: `bar`,
		})
		return
	}
	// TODO: insert queue entry, input: priority, name
}
func UpdateQueue(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_update-queue.html`, gin.H{})
		return
	}
	// TODO: update queue entry, input: id, priority, name
}
func RemoveQueue(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_remove-queue.html`, gin.H{})
		return
	}
	// TODO: remove top priority entry, input: limit=1
}
func ShowQueue(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_show-queue.html`, gin.H{})
		return
	}
	// TODO: show queue entries (order by priority), input: limit=5
}
