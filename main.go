package main

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/admin"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/chat"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/server"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user"
)


func main () {
	cnfg,err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading configuration file")
	}

	server := server.NewServer()
	server.R.LoadHTMLGlob("templates/*")
	user.NewUserRoute(server.R,*cnfg)
	admin.NewAdminRoute(server.R,*cnfg)
	chat.NewChatRoutes(server.R,*cnfg)
	server.StartServer(cnfg.APIPORT)
}