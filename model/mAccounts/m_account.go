package mAccounts

import (
	"errors"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Number string
	Address string
	Phone string
}

func (a *Account) IsValid() error {
	if a.Number == `` {
		return errors.New(`number may not be empty`)
	}
	if a.Address == `` {
		return errors.New(`address may not be emmpty`)
	} 
	if a.Phone == `` {
		return errors.New(`phone may not be empty`)
	}
	return nil
}
