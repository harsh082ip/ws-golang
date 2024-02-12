package router

import (
	// ""

	"ws-go/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitialRouter(wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {

	return r.Run(addr)
}
