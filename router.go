package mdview

import (
	"github.com/gin-gonic/gin"
	"github.com/go-steven/mdview/handler"
	"log"
	"os"
	"path"
	"runtime"
)

func Router(r *gin.Engine) {
	r.Delims("{{", "}}")
	r.LoadHTMLGlob(curr_path(1) + "/templates/*.tpl")

	handler.SetLogger(logger)
	g := r.Group("/mdview")
	{
		g.GET("/list/*path", handler.IndexHandler)
		g.GET("/md/*doc", handler.DocHandler)
	}
}

var logger *log.Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

func SetLogger(l *log.Logger) {
	logger = l
}

func curr_path(skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	return path.Dir(file)
}
