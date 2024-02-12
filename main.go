package main

import (
	"ws-go/router"
	"ws-go/ws"
)

func main() {

	hub := ws.NewHub()
	wsHanler := ws.NewHandler(hub)
	go hub.Run()

	router.InitialRouter(wsHanler)
	router.Start(":8083")
}
