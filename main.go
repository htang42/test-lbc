package main

import (
	"github.com/gin-gonic/gin"
	"github.com/htang42/test-lbc/handler/fizzbuzz"
)

func main() {
	r := gin.Default()

	fizzbuzz.RegisterRoutes(r)

	r.Run()
}
