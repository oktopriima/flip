/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21:45
**/

package models

import "time"

type Withdrawals struct {
	ID              int       `json:"id"`
	TransactionCode int64     `json:"transaction_code"`
	BankCode        string    `json:"bank_code"`
	AccountNumber   string    `json:"account_number"`
	Amount          float64   `json:"amount"`
	Remark          string    `json:"remark"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
