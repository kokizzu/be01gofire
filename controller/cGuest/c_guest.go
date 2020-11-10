package cGuest

import (
	"be01gofire/controller"
	"be01gofire/model/mBank"
	"be01gofire/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func Login(ctx *controller.Ctx) {
//	if ctx.Context.Request.Method == `GET` {
//		ctx.HTML(http.StatusOK, `guest_login.html`, gin.H{
//			`foo`: `bar`,
//		})
//		return
//	}
//	user := mUser.User{}
//	err := ctx.ShouldBindJSON(&user)
//	pass := user.Pass
//	res := map[string]interface{}{}
//	if err == nil {
//		err := user.FindByEmail(ctx.Firestore)
//		if err == nil {
//			if user.CheckPass(pass) {
//				res[`email`] = user.Email
//				// TODO: set cookie/jwt
//			} else {
//				res[`error`] = `wrong username or password` // wrong password
//			}
//		} else {
//			res[`error`] = `wrong username or password` // not exists
//		}
//	}
//	if err != nil {
//		res[`error`] = err.Error()
//	}
//	ctx.JSON(http.StatusOK, res)
//}
//
//func Register(ctx *controller.Ctx) {
//	if ctx.Context.Request.Method == `GET` {
//		ctx.HTML(http.StatusOK, `guest_register.html`, gin.H{
//			`title`:`Register ` + time.Now().Format(time.RFC1123Z), 
//		})
//		return
//	}
//	// https://github.com/gin-gonic/gin#bind-form-data-request-with-custom-struct 
//	user := mUser.User{}
//	err := ctx.ShouldBindJSON(&user)
//	res := map[string]interface{}{}
//	if err == nil {
//		user.HashPassword()
//		err = user.Insert(ctx.Firestore)
//		if err == nil {
//			res[`user`] = user
//		}
//	}
//	if err != nil {
//		res[`error`] = err.Error()
//	}
//	ctx.JSON(http.StatusOK, res)
//}
//
//func AddQueue(ctx *controller.Ctx) {
//	if ctx.Context.Request.Method == `GET` {
//		ctx.HTML(http.StatusOK, `guest_add-queue.html`, gin.H{})
//		return
//	}
//	qe := mQueue.QueueEntry{}
//	err := ctx.ShouldBindJSON(&qe)
//	res := map[string]interface{}{}
//	if err == nil {
//		err = qe.Insert(ctx.Firestore)
//		if err == nil {
//			res[`entry`] = qe
//		}
//	}
//	if err != nil {
//		res[`error`] = err.Error()
//	}
//	ctx.JSON(http.StatusOK, res)	
//}
//func UpdateQueue(ctx *controller.Ctx) {
//	if ctx.Context.Request.Method == `GET` {
//		ctx.HTML(http.StatusOK, `guest_update-queue.html`, gin.H{})
//		return
//	}
//	qe := mQueue.QueueEntry{}
//	err := ctx.ShouldBindJSON(&qe)
//	res := map[string]interface{}{}
//	if err == nil {
//		err = qe.Update(ctx.Firestore)
//		if err == nil {
//			res[`entry`] = qe
//		} else {
//			res[`error`] = `entry not found`
//		}
//	}
//	if err != nil {
//		res[`error`] = err.Error()
//	}
//	ctx.JSON(http.StatusOK, res)
//}
//func RemoveQueue(ctx *controller.Ctx) {
//	if ctx.Context.Request.Method == `GET` {
//		ctx.HTML(http.StatusOK, `guest_remove-queue.html`, gin.H{})
//		return
//	}
//	qe := mQueue.QueueEntry{}
//	err := ctx.ShouldBindJSON(&qe)
//	res := map[string]interface{}{}
//	if err == nil {
//		err = qe.Delete(ctx.Firestore)
//		if err == nil {
//			res[`entry`] = qe
//		} else {
//			res[`error`] = `entry not found`
//		}
//	}
//	if err != nil {
//		res[`error`] = err.Error()
//	}
//	ctx.JSON(http.StatusOK, res)
//	
//}
//
//type ShowQueueInput struct {
//	First int
//}
//
//func ShowQueue(ctx *controller.Ctx) {
//	if ctx.Context.Request.Method == `GET` {
//		ctx.HTML(http.StatusOK, `guest_show-queue.html`, gin.H{})
//		return
//	}
//	qe := mQueue.QueueEntry{}
//	sqi := ShowQueueInput{}
//	err := ctx.ShouldBindJSON(&sqi)
//	res := map[string]interface{}{}
//	rows, err := qe.List(ctx.Firestore,sqi.First)
//	if err != nil {
//		res[`error`] = err.Error()
//	}
//	res[`list`] = rows
//	ctx.JSON(http.StatusOK, res)
//}

//  curl localhost:8084/guest/create-account -H 'content-type:application/json' -d '{"email":"kiswono@gmail.com","id":1}'
//{"account":{"ID":1,"id_account":"id-414","name":"","email":"kiswono@gmail.com","password":"$2a$10$SWJcgx9kOWlPSJVRk6r9S.40eirOUKlYKr2P8LBKid6.uFFWHNqZi","account_number":832712,"saldo":0}}%   
func CreateAccount(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_create-account.html`, gin.H{})
		return
	}
	a := mBank.Account{}
	err := ctx.ShouldBindJSON(&a)
	res := map[string]interface{}{}
	if err == nil {
		pass := utils.HashGenerator(a.Password) 
		a.Password = pass
		err = a.InsertNewAccount(ctx.Db)
		if err == nil {
			res[`account`] = a
		} 
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

// curl localhost:8084/guest/login -H 'content-type:application/json' -d '{"email":"kiswono@gmail.com","password":""}'
// {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X251bWJlciI6ODMyNzEyLCJlbWFpbCI6Imtpc3dvbm9AZ21haWwuY29tIn0.VrIftkT0vfYxDWmjvuK_aDiIcevzVitiYtr9QqAMclQ"}%
func Login(ctx *controller.Ctx) {
	auth := mBank.Auth{}
	err := ctx.ShouldBindJSON(&auth)
	res := map[string]interface{}{}
	if err == nil {
		token := ``
		err, token = auth.Login(ctx.Db)
		if err == nil {
			res[`token`] = token
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK,res)
}
