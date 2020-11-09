package main

import (
	"encoding/json"
	"fmt"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"io"
	"net/http"
	"strings"
	"testing"
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
	c := &http.Client{}
	// HIT API register dengan username random password random, check tidak boleh error
	email := `dummy.` + S.RandomPassword(12) + `@gmail.com`
	pass := S.RandomPassword(5)
	json := M.SX{
		`email`: email,
		`pass`:  pass,
	}
	body := strings.NewReader(map2json(json))
	req, err := http.NewRequest(`POST`, `http://localhost:8084/guest/create-account`, body)
	if err != nil {
		t.Error(err)
	}
	res, err := c.Do(req)
	if err != nil {
		t.Error(err)
	}
	json = json2map(res.Body)
	acc := json.GetMSX(`account`)
	email2 := acc.GetStr(`email`)
	if email != email2 {
		t.Errorf(`expecting email: %s but got %s instead`, email, email2)
	}
	// HIT API register dengan username yg sama lagi, check errornya
	// HIT API login dengan username benar, password salah, check harus error
	// HIT API login dengan username benar, password benar, check harus dapat token
	// HIT API deposit dengan token salah, check harus error
	// HIT API check saldo dengan token benar, check saldo mula2
	// HIT API deposit dengan token benar, check tidak boleh error
	// HIT API check saldo dengan token benar, check saldo harus bertambah
	// HIT API deposit dengan token benar, check tidak boleh error
	// HIT API check saldo dengan token benar, check saldo harus bertambah lagi dengan jumlah benar
}
