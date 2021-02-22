/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:17
**/

package main

import (
	"github.com/oktopriima/flip/container"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	c := dig.New()

	c = container.NewBuildConfigContainer(c)
	c = container.NewControllerContainer(c)

	return c
}
