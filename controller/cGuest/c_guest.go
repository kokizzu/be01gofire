package cGuest

import (
	"be01gofire/controller"
	"be01gofire/model/mAccounts"
	"be01gofire/model/mQueue"
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
		ctx.HTML(http.StatusOK, `guest_add-queue.html`, gin.H{})
		return
	}
	qe := mQueue.QueueEntry{}
	err := ctx.ShouldBindJSON(&qe)
	res := map[string]interface{}{}
	if err == nil {
		err = qe.Insert(ctx.Firestore)
		if err == nil {
			res[`entry`] = qe
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)	
}
func UpdateQueue(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_update-queue.html`, gin.H{})
		return
	}
	qe := mQueue.QueueEntry{}
	err := ctx.ShouldBindJSON(&qe)
	res := map[string]interface{}{}
	if err == nil {
		err = qe.Update(ctx.Firestore)
		if err == nil {
			res[`entry`] = qe
		} else {
			res[`error`] = `entry not found`
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}
func RemoveQueue(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_remove-queue.html`, gin.H{})
		return
	}
	qe := mQueue.QueueEntry{}
	err := ctx.ShouldBindJSON(&qe)
	res := map[string]interface{}{}
	if err == nil {
		err = qe.Delete(ctx.Firestore)
		if err == nil {
			res[`entry`] = qe
		} else {
			res[`error`] = `entry not found`
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
	
}

type ShowQueueInput struct {
	First int
}

func ShowQueue(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_show-queue.html`, gin.H{})
		return
	}
	qe := mQueue.QueueEntry{}
	sqi := ShowQueueInput{}
	err := ctx.ShouldBindJSON(&sqi)
	res := map[string]interface{}{}
	rows, err := qe.List(ctx.Firestore,sqi.First)
	if err != nil {
		res[`error`] = err.Error()
	}
	res[`list`] = rows
	ctx.JSON(http.StatusOK, res)
}

//  curl -X POST localhost:8084/guest/create-account -H 'content-type:application/json' -d '{"number":"124-125-1265","address":"gayungsari 3"}'
// {"error":"phone may not be empty"}
// curl -X POST localhost:8084/guest/create-account -H 'content-type:application/json' -d '{"number":"124-125-1265","address":"gayungsari 3","phone":"9812958125"}'
// {"account":{"ID":1,"CreatedAt":"2020-11-03T09:00:07.48785323+07:00","UpdatedAt":"2020-11-03T09:00:07.48785323+07:00","DeletedAt":"0001-01-01T00:00:00Z","Number":"124-125-1265","Address":"gayungsari 3","Phone":"9812958125"}}
func CreateAccount(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_create-account.html`, gin.H{})
		return
	}
	a := mAccounts.Account{}
	err := ctx.ShouldBindJSON(&a)
	res := map[string]interface{}{}
	if err == nil {
		err = a.IsValid()
	}
	if err == nil {
		tx := ctx.Db.Create(&a) // INSERT INTO ... VALUES ...
		err = tx.Error
		if err == nil {
			res[`account`] = a
		} 
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}
