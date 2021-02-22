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
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/flip/helper/configurations"
	sbig_api "github.com/oktopriima/flip/helper/sbig-api"
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

	// register application configuration
	if err = container.Provide(func() configurations.Config {
		return configurations.NewConfig()
	}); err != nil {
		panic(err)
	}

	// register database configuration
	if err = container.Provide(func(conf configurations.Config) (*gorm.DB, error) {
		return configurations.NewMysqlClient(conf.GetStringMap("mysql"))
	}); err != nil {
		panic(err)
	}

	// register S BIG API
	if err = container.Provide(func(config configurations.Config) (*sbig_api.Client, error) {
		return sbig_api.NewClient(config)
	}); err != nil {
		panic(err)
	}

	return container
}
