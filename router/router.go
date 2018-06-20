package router

import (
	"github.com/zhs007/tradingview_udf_go/controller/index"

	"github.com/gin-gonic/gin"
)

// Router -
var Router *gin.Engine

func init() {
	Router = gin.Default()
}

// SetRouter -
func SetRouter() {
	Router.GET("/", index.Index())
}
