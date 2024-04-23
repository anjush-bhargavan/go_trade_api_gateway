package admin

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_api_gateway/middleware"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/admin/adminpb"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

// Admin represents the admin route handler, containing configuration and gRPC client.
type Admin struct {
	Cfg    *config.Config
	Client pb.AdminServiceClient
}

// NewAdminRoute initializes the admin routes and handlers.
func NewAdminRoute(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	adminHandler := &Admin{
		Cfg:    &cfg,
		Client: client,
	}

	apiVersion := c.Group("/api/v1")

	admin := apiVersion.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)
		
	}

	auth := admin.Group("/auth")
	auth.Use(middleware.AdminAuthorization(cfg.SECRETKEY,"admin"))
	{
		auth.POST("/category",adminHandler.AddCategory)
		auth.PATCH("/category/:id",adminHandler.EditCategory)
		auth.DELETE("/category/:id",adminHandler.RemoveCategory)
		auth.GET("/category/:id",adminHandler.FindCategory)
		auth.GET("/category",adminHandler.FindCategories)

		auth.DELETE("/product/:id",adminHandler.RemoveProduct)
		auth.GET("/product/:id",adminHandler.FindProduct)
		auth.GET("/products",adminHandler.FindProducts)

		auth.PATCH("/user/:id",adminHandler.BlockUser)

		auth.GET("/orders",adminHandler.OrderHistory)
		auth.GET("/orders/:id",adminHandler.UserOrders)
		auth.GET("/order/:id",adminHandler.FindOrder)
	}
}
