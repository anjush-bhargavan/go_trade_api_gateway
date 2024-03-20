package server

import "github.com/gin-gonic/gin"

// SeverStruct represent the model of server it will have gin engine
type ServerStruct struct {
	R *gin.Engine
}

// StartServer method starts the server in the specified port
func (s *ServerStruct) StartServer(port string) {
	s.R.Run(":" + port)
}

// Sever method will return the SeverStruct instance with default gin engine attached
func Server() *ServerStruct {
	engine := gin.Default()

	return &ServerStruct{
		R: engine,
	}
}
