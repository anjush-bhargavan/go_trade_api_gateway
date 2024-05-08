package chat

import (
	"log"

	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/chat/pb"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user"
	userpb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// Chat represents the chat route handler, containing configuration and gRPC client.
type Chat struct {
	cfg        *config.Config
	userClient userpb.UserServiceClient
	client     pb.ChatServiceClient
}

// NewChatRoutes initializes the chat routes and handlers.
func NewChatRoutes(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	userClient, err := user.ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc user client : %v", err.Error())
	}

	chatHandler := &Chat{
		cfg:        &cfg,
		client:     client,
		userClient: userClient,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/user")
	{
		user.GET("/chat", chatHandler.Chat)
	}
	c.GET("/chat", chatHandler.ChatPage)
}
