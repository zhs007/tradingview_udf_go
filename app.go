package main

import (
	"./router"
)

// StartServ -
func StartServ(servaddr string) {
	r := router.Router
	router.SetRouter()
	r.Run(servaddr)
}
