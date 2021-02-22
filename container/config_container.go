/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:16
**/

package container

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func NewBuildConfigContainer(container *dig.Container) *dig.Container {
	var err error

	// register gin engine
	if err = container.Provide(func() *gin.Engine {
		return gin.Default()
	}); err != nil {
		panic(err)
	}

	return container
}
