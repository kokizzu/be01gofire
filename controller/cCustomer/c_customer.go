package cCustomer

import (
	"be01gofire/controller"
	"be01gofire/model/mBank"
	"fmt"
	"net/http"
)

func checkLogin(ctx *controller.Ctx) (map[string]interface{}, int, bool) {
	res := map[string]interface{}{}
	idAccount := controller.CheckAuth(ctx.Context)
	if idAccount < 0 {
		res[`error`] = `not yet logged in`
		ctx.JSON(http.StatusOK, res)
		return nil, 0, true
	}
	return res, idAccount, false
}


// curl localhost:8084/customer/account -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X251bWJlciI6ODMyNzEyLCJlbWFpbCI6Imtpc3dvbm9AZ21haWwuY29tIn0.VrIftkT0vfYxDWmjvuK_aDiIcevzVitiYtr9QqAMclQ'
// {"account":{"ID":1,"id_account":"id-414","name":"","email":"kiswono@gmail.com","password":"$2a$10$SWJcgx9kOWlPSJVRk6r9S.40eirOUKlYKr2P8LBKid6.uFFWHNqZi","account_number":832712,"saldo":0},"transaction":[]}
func Account(ctx *controller.Ctx) {
	res, idAccount, done := checkLogin(ctx)
	if done {
		return
	}
	a := mBank.Account{}
	a.AccountNumber = idAccount
	err, transactions := a.GetAccountDetail(ctx.Db)
	fmt.Printf("%#v\n",a)
	if err == nil {
		res[`account`] = a
		res[`transaction`] = transactions
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK,res)
}

func Transfer(ctx *controller.Ctx) {
	res, idAccount, done := checkLogin(ctx)
	if done {
		return
	}
	transaction := mBank.Transaction{}
	err := ctx.ShouldBindJSON(&transaction)
	transaction.Sender = idAccount
	if err == nil {
		err = transaction.Transfer(ctx.Db)
		if err == nil {
			res[`transaction`] = transaction
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK,res)
}

func Widthdraw(ctx *controller.Ctx) {
	res, idAccount, done := checkLogin(ctx)
	if done {
		return
	}
	transaction := mBank.Transaction{}
	err := ctx.ShouldBindJSON(&transaction)
	transaction.Sender = idAccount
	if err == nil {
		err = transaction.Withdraw(ctx.Db)
		if err == nil {
			res[`transaction`] = transaction
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

//  curl localhost:8084/customer/deposit -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X251bWJlciI6ODMyNzEyLCJlbWFpbCI6Imtpc3dvbm9AZ21haWwuY29tIn0.VrIftkT0vfYxDWmjvuK_aDiIcevzVitiYtr9QqAMclQ' -H 'content-type:application/json' -d '{"transaction_description":"ATM Bersama 0125125","amount":50000}' 
//{"transaction":{"transaction_type":2,"transaction_description":"ATM Bersama 0125125","sender":832712,"amount":50000,"recipient":0,"timestamp":1604457608}}
func Deposit(ctx *controller.Ctx) {
	res, idAccount, done := checkLogin(ctx)
	if done {
		return
	}
	transaction := mBank.Transaction{}
	err := ctx.ShouldBindJSON(&transaction)
	transaction.Sender = idAccount
	fmt.Printf("%#v\n",transaction)
	if err == nil {
		err = transaction.Deposit(ctx.Db)
		if err == nil {
			res[`transaction`] = transaction
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK,res)
}
