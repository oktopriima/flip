/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:21
**/

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/flip/database/services"
	sbig_api "github.com/oktopriima/flip/helper/sbig-api"
	"github.com/oktopriima/flip/request"
	"net/http"
)

type DisbursementController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type disbursementController struct {
	slightlyBigClient  *sbig_api.Client
	withdrawalServices services.WithdrawalServices
}

func (d *disbursementController) Create(ctx *gin.Context) {
	var req request.DisbursementRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Redirect(http.StatusOK, "disbursement")
	}
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

	fmt.Println(resp)

	ctx.HTML(http.StatusOK, "disbursement/index.tmpl", nil)
}

func (d *disbursementController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "disbursement/index.tmpl", nil)
}

func NewDisbursementController(slightlyBigClient *sbig_api.Client,
	withdrawalServices services.WithdrawalServices) DisbursementController {
	return &disbursementController{
		slightlyBigClient:  slightlyBigClient,
		withdrawalServices: withdrawalServices,
	}
}
