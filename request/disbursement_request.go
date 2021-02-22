/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:35
**/

package request

// request body from front end or mobile
type DisbursementRequest struct {
	BankCode      string  `form:"bank_code"`
	AccountNumber string  `form:"account_number"`
	Amount        float64 `form:"amount"`
	Remark        string  `form:"remark"`
}
