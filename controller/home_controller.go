/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:29
**/

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController interface {
	Index(ctx *gin.Context)
}

type indexController struct {
}

func (i *indexController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "homepage/index.tmpl", nil)
}

func NewIndexController() HomeController {
	return &indexController{}
}
