/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:11
**/

package main

import (
	"github.com/oktopriima/flip/routes"
	"log"
)

func main() {

	c := NewContainer()

	if err := c.Invoke(routes.InvokeWeb); err != nil {
		log.Fatalf("failed build application %v", err)
	}

	if err := c.Provide(NewServer); err != nil {
		log.Fatalf("failed provide server. error %v", err)
	}

	if err := c.Invoke(func(s *Server) {
		s.Run()
	}); err != nil {
		log.Fatal(err)
	}
}
