/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22:19
**/

package sbig_api

type Response struct {
	ID              int64   `json:"id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	Timestamp       string  `json:"timestamp"`
	BankCode        string  `json:"bank_code"`
	AccountNumber   string  `json:"account_number"`
	BeneficiaryName string  `json:"beneficiary_name"`
	Remark          string  `json:"remark"`
	Receipt         string  `json:"receipt"`
	TimeServed      string  `json:"time_served"`
	Fee             float64 `json:"fee"`
}

type CreateDisbursementRequest struct {
	BankCode      string  `url:"bank_code"`
	AccountNumber string  `url:"account_number"`
	Amount        float64 `url:"amount"`
	Remark        string  `url:"remark"`
}

type FindDisbursementRequest struct {
	TransactionID int64 `url:"transaction_id"`
}
