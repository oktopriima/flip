/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:30
**/

package container

import (
	"github.com/oktopriima/flip/controller"
	"go.uber.org/dig"
)

func NewControllerContainer(container *dig.Container) *dig.Container {
	var err error
	if err = container.Provide(controller.NewIndexController); err != nil {
		panic(err)
	}

	if err = container.Provide(controller.NewDisbursementController); err != nil {
		panic(err)
	}

	return container
}
