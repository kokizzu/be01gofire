package mBank

import (
	"be01gofire/utils"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

const (
	TRANSFER = 0
	WITHDRAW = 1
	DEPOSIT  = 2
)

type Account struct {
	ID            int    `gorm:"primary_key", json:"-"`
	IdAccount     string `json:"id_account,omitempty"`
	Name          string `json:"name"`
	Email         string `gorm:"email,unique" json:"email"`
	Password      string `json:"password,omitempty"`
	AccountNumber int    `json:"account_number,omitempty"` // TODO: set jadi unique
	Saldo         int64  `json:"saldo"`
}

func (account *Account) InsertNewAccount(db *gorm.DB) error {
	account.AccountNumber = utils.RangeIn(111111, 999999)
	account.Saldo = 0
	account.IdAccount = fmt.Sprintf("id-%d", utils.RangeIn(111, 999))
	if err := db.Create(&account).Error; err != nil {
		return fmt.Errorf("invalid prepare statement :%+v\n", err)
	}
	return nil
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (auth *Auth) Login(db *gorm.DB) (error, string) {
	account := Account{}
	if strings.TrimSpace(auth.Email) == `` {
		return errors.New(`email may not be empty`), ``
	}
	if err := db.Where(&Account{Email: auth.Email}).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("account not found"), ""
		}
	}
	if !utils.SamePassword(account.Password,auth.Password) {
		return errors.New("incorrect Password"), ""
	} else {
		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": auth.Email, "account_number": account.AccountNumber})
		token, err := sign.SignedString([]byte("secret"))
		if err != nil {
			return err, ""
		}
		return nil, token
	}
}

type Transaction struct {
	ID                     int    `gorm:"primary_key" json:"-"`
	TransactionType        int    `json:"transaction_type,omitempty"`
	TransactionDescription string `json:"transaction_description"`
	Sender                 int    `json:"sender"`
	Amount                 int64  `json:"amount"`
	Recipient              int    `json:"recipient"`
	Timestamp              int64  `json:"timestamp,omitempty"`
}

func (account *Account) GetAccountDetail(db *gorm.DB) (error, []Transaction) {
	var transaction []Transaction
	idAccount := account.AccountNumber
	res := *account
	if err := db.Where("sender = ? OR recipient = ?", idAccount, idAccount).Find(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("account not found"), transaction
		} else {
			return fmt.Errorf("invalid prepare statement :%+v\n", err), transaction
		}
	}
	if err := db.Where(&Account{AccountNumber: idAccount}).Find(&res).Error; err != nil {
		return fmt.Errorf("error fetching account :%+v\n", err), transaction
	}
	*account = res
	return nil, transaction
}

func (transaction *Transaction) Transfer(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		var sender, recipient Account
		if err := tx.Model(&Account{}).Where(&Account{AccountNumber: transaction.Sender}).First(&sender).Update(`saldo`, sender.Saldo-transaction.Amount).Error; err != nil {
			return err
		}
		if err := tx.Model(&Account{}).Where(&Account{AccountNumber: transaction.Recipient}).First(&recipient).Update(`saldo`, recipient.Saldo+transaction.Amount).Error; err != nil {
			log.Println("ERROR : " + err.Error())
			return err
		}
		transaction.TransactionType = TRANSFER
		transaction.Timestamp = time.Now().Unix()
		if err := tx.Create(&transaction).Error; err != nil {
		}
		return nil
	})
	return err
}

func (transaction *Transaction) Withdraw(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		var sender Account
		if err := tx.Model(&Account{}).Where(&Account{AccountNumber: transaction.Sender}).First(&sender).Update("saldo", sender.Saldo-transaction.Amount).Error; err != nil {
			return err
		}
		transaction.TransactionType = WITHDRAW
		transaction.Timestamp = time.Now().Unix()
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (transaction *Transaction) Deposit(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		var sender Account
		if err := tx.Model(&Account{}).Where(&Account{AccountNumber: transaction.Sender}).First(&sender).Update("saldo", sender.Saldo+transaction.Amount).Error; err != nil {
			return err
		}
		transaction.TransactionType = DEPOSIT
		transaction.Timestamp = time.Now().Unix()
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
