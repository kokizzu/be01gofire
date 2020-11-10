package main

import (
	"encoding/json"
	"fmt"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"
)

func map2json(m M.SX) string {
	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println(`ERR map2json: ` + err.Error())
	}
	return string(json)
}

func json2map(r io.Reader) M.SX {
	res := M.SX{}
	err := json.NewDecoder(r).Decode(&res)
	if err != nil {
		fmt.Println(`ERR json2map: ` + err.Error())
	}
	return res
}

func TestRegisterLoginDepositCheckSaldoApi(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	c := &http.Client{}
	// HIT API register dengan username random password random, check tidak boleh error
	email := `dummy.` + S.RandomPassword(12) + `@gmail.com`
	pass := S.RandomPassword(5)
	json := M.SX{
		`email`: email,
		`password`:  pass,
	}
	json = hitApi(t, `/guest/create-account`, json, c, nil)
	acc := json.GetMSX(`account`)
	expect(t, acc.GetStr(`email`), email)
	// HIT API register dengan username yg sama lagi, check errornya
	json = M.SX{
		`email`: email,
		`password`: pass,
	}
	json = hitApi(t, `/guest/create-account`, json, c, nil)
	expect(t, json.GetStr(`error`), `Duplicate`)
	// HIT API dengan email kosong
	json = M.SX{}
	json = hitApi(t, `/guest/login`, json, c, nil)
	expect(t,json.GetStr(`error`),`empty`)
	// HIT API login dengan username benar, password salah, check harus error
	json = M.SX{
		`email`: email,
	}
	json = hitApi(t, `/guest/login`, json, c, nil)
	expect(t,json.GetStr(`error`),`incorrect`)
	// HIT API login dengan username benar, password benar, check harus dapat token
	json = M.SX{
		`email`:    email,
		`password`: pass,
	}
	json = hitApi(t, `/guest/login`, json, c, nil)
	token := json.GetStr(`token`)
	if token == `` {
		t.Errorf(`expecting token, but got %v instead`, json)
	}
	// HIT API deposit dengan token salah, check harus error
	json = M.SX{}
	json = hitApi(t, `/customer/deposit`, json, c, M.SS{
		`Authorization`: `j3890v475630568`,
	})
	expect(t,json.GetStr(`error`),`invalid`)
	// HIT API check saldo dengan token benar, check saldo mula2
	json = M.SX{}
	json = hitApi(t,`/customer/account`, json, c, M.SS{
		`Authorization`: token,
	})
	acc = json.GetMSX(`account`)
	expect(t,acc.GetStr(`email`),email) 
	saldo := json.GetInt(`saldo`)
	_ = saldo
	// HIT API deposit dengan token benar, check tidak boleh error
	// HIT API check saldo dengan token benar, check saldo harus bertambah
	// HIT API deposit dengan token benar, check tidak boleh error
	// HIT API check saldo dengan token benar, check saldo harus bertambah lagi dengan jumlah benar
}

func expect(t *testing.T, errStr string, expStr string) {
	if !strings.Contains(errStr, expStr) {
		t.Errorf(`expecting error: %s, but got %v instead`, expStr, errStr)
	}
}

func hitApi(t *testing.T, url string, json M.SX, c *http.Client, headers M.SS) M.SX {
	fmt.Printf("Hitting API: %s %v\n",url,json)
	body := strings.NewReader(map2json(json))
	req, err := http.NewRequest(`POST`, `http://localhost:8084`+url, body)
	if err != nil {
		t.Error(err)
	}
	for k, v := range headers {
		req.Header.Set(k,v)
	}
	res, err := c.Do(req)
	if err != nil {
		t.Error(err)
	}
	json = json2map(res.Body)
	fmt.Printf("%#v\n", json)
	return json
}
