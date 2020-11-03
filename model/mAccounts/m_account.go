package mAccounts

import (
	"errors"
	"gorm.io/gorm"
)

// 1. pastikan bisa connect ke mysql lokal (controller/server.go --> MysqlDsn diubah)
// 2. buat route untuk masing2 API
// 3. buat model: copy2 dari bank.go, tambahkan fungsi yg dibutuhkan
//        dan dari bagian Melengkapi model
// 4. lalu panggil2 dari controller
// 5. buat view/curl, lalu testing

type Account struct {
	gorm.Model
	Number string
	Address string
	Phone string
}
type Transaction struct {
	
	AccountId int 
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
