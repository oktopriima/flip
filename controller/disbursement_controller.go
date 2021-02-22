/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:21
**/

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/flip/database/models"
	"github.com/oktopriima/flip/database/services"
	sbig_api "github.com/oktopriima/flip/helper/sbig-api"
	"github.com/oktopriima/flip/request"
	"net/http"
)

type DisbursementController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Find(ctx *gin.Context)
}

type disbursementController struct {
	db                 *gorm.DB
	slightlyBigClient  *sbig_api.Client
	withdrawalServices services.WithdrawalServices
}

func (d *disbursementController) Find(ctx *gin.Context) {
	var req request.DisbursementSearchRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Redirect(http.StatusOK, "disbursement")
	}

	data, err := d.withdrawalServices.FindOne(map[string]interface{}{
		"transaction_code": req.TransactionID,
	})
	if err != nil {
		ctx.HTML(http.StatusOK, "disbursement/index.tmpl", nil)
		return
	}

	resp, err := d.slightlyBigClient.FindDisbursement(sbig_api.FindDisbursementRequest{TransactionID: data.TransactionCode})
	if err != nil {
		ctx.HTML(http.StatusOK, "disbursement/index.tmpl", nil)
		return
	}

	data.Receipt = resp.Receipt
	data.Status = resp.Status
	data.TimeServed = resp.TimeServed

	tx := d.db.Begin()
	defer tx.Rollback()

	if err := d.withdrawalServices.Update(data, tx); err != nil {
		ctx.HTML(http.StatusOK, "disbursement/index.tmpl", nil)
		return
	}

	tx.Commit()
	ctx.HTML(http.StatusOK, "disbursement/detail.tmpl", gin.H{
		"transaction_code": data.TransactionCode,
		"amount":           data.Amount,
		"bank_code":        data.BankCode,
		"account_number":   data.AccountNumber,
		"status":           data.Status,
		"receipt":          data.Receipt,
		"time_served":      data.TimeServed,
	})
}

func (d *disbursementController) Create(ctx *gin.Context) {
	var req request.DisbursementRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Redirect(http.StatusOK, "disbursement")
	}

	// create http request to 3rd party API
	resp, err := d.slightlyBigClient.CreateDisbursement(sbig_api.CreateDisbursementRequest{
		BankCode:      req.BankCode,
		AccountNumber: req.AccountNumber,
		Amount:        req.Amount,
		Remark:        req.Remark,
	})
	if err != nil {
		ctx.Redirect(http.StatusOK, "disbursement")
		return
	}

	tx := d.db.Begin()
	defer tx.Rollback()

	data, err := d.withdrawalServices.Insert(&models.Withdrawals{
		TransactionCode: resp.ID,
		BankCode:        resp.BankCode,
		AccountNumber:   resp.AccountNumber,
		Amount:          resp.Amount,
		Remark:          resp.Remark,
		Status:          resp.Status,
		Receipt:         resp.Receipt,
		TimeServed:      resp.TimeServed,
	}, tx)

	// just redirect if theres an error when save data into database
	if err != nil {
		ctx.Redirect(http.StatusOK, "disbursement")
		return
	}

	tx.Commit()

	ctx.HTML(http.StatusOK, "disbursement/detail.tmpl", gin.H{
		"transaction_code": data.TransactionCode,
		"amount":           data.Amount,
		"bank_code":        data.BankCode,
		"account_number":   data.AccountNumber,
		"status":           data.Status,
		"receipt":          data.Receipt,
		"time_served":      data.TimeServed,
	})
}

func (d *disbursementController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "disbursement/index.tmpl", nil)
}

func NewDisbursementController(db *gorm.DB,
	slightlyBigClient *sbig_api.Client,
	withdrawalServices services.WithdrawalServices) DisbursementController {
	return &disbursementController{
		db:                 db,
		slightlyBigClient:  slightlyBigClient,
		withdrawalServices: withdrawalServices,
	}
}
