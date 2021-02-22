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
	"net/http"
)

type DisbursementController interface {
	Index(ctx *gin.Context)
}

type disbursementController struct {
}

func (d *disbursementController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "disbursement/index.tmpl", nil)
}

func NewDisbursementController() DisbursementController {
	return &disbursementController{}
}
