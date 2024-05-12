package app

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	registerHandlerFunc(r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
