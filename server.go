/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:21
**/

package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(engine *gin.Engine) *Server {
	return &Server{engine: engine}
}

func (s *Server) Run() {
	if err := s.engine.Run(":5000"); err != nil {
		log.Fatalf("failed run application. error %v", err)
	}
}
