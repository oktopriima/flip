/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:12
**/

package route

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/flip/controller"
)

func InvokeWeb(engine *gin.Engine,
	home controller.HomeController,
	disbursement controller.DisbursementController,
) {

	engine.LoadHTMLGlob("templates/**/*")

	route := engine.Group("/")

	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(gin.ErrorLogger())

	route.GET("/", home.Index)

	// disburse route group
	{
		disburseRoute := route.Group("disbursement")
		disburseRoute.GET("", disbursement.Index)
	}
}
