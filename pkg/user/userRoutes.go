package user

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_api_gateway/middleware"
	chatpb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/chat/pb"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// User represents the User route handler, containing configuration and gRPC client.
type User struct {
	cfg        *config.Config
	client     pb.UserServiceClient
	chatClient chatpb.ChatServiceClient
}

// NewUserRoute initializes the user routes and handlers.
func NewUserRoute(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	userHandler := &User{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/user")

	{
		user.POST("/signup", userHandler.UserSignup)
		user.POST("/verify", userHandler.UserVerify)
		user.POST("/login", userHandler.UserLogin)
	}
	auth := user.Group("/auth")
	auth.Use(middleware.Authorization(cfg.SECRETKEY))
	{
		auth.POST("/address", userHandler.AddAddress)
		auth.PATCH("/address/:id", userHandler.EditAddress)
		auth.DELETE("/address/:id", userHandler.RemoveAddress)

		auth.GET("/profile", userHandler.ViewProfile)
		auth.PATCH("/profile", userHandler.EditProfile)
		auth.PATCH("/password", userHandler.ChangePassword)

		// auth.POST("/seller", userHandler.BeASeller)
		auth.POST("/product", userHandler.AddProduct)
		auth.PATCH("/product/:id", userHandler.EditProduct)
		auth.DELETE("/product/:id", userHandler.RemoveProduct)

		auth.GET("/product/:id", userHandler.FindProduct)
		auth.GET("/products", userHandler.FindAllProducts)
		auth.GET("/category/product/:id", userHandler.FindProductsByCategory)
		auth.GET("/category/:id", userHandler.FindCategory)
		auth.GET("/categories", userHandler.FindAllCategories)

		auth.POST("/watchlist/:id", userHandler.AddToWatchlist)
		auth.GET("/watchlist", userHandler.ViewWatchlist)

		auth.POST("/bid", userHandler.AddBid)
		auth.GET("/bids/:id", userHandler.GetBids)

		user.GET("/payment", userHandler.Payment)
		user.GET("/payment/success", userHandler.PaymentSuccess)
		user.GET("/success", userHandler.SuccessPage)
	}
}
