package main

import (
	"github.com/zhs007/tradingview_udf_go/router"
)

// StartServ -
func StartServ(servaddr string) {
	r := router.Router
	router.SetRouter()
	r.Run(servaddr)
}
