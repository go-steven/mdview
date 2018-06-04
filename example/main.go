package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-steven/mdview"
)

func main() {
	r := gin.New()
	mdview.Router(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
