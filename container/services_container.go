/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:47
**/

package container

import (
	"github.com/oktopriima/flip/database/services"
	"go.uber.org/dig"
)

func NewServicesContainer(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(services.NewWithdrawalServices); err != nil {
		panic(err)
	}

	return container
}
