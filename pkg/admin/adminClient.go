package admin

import (
	"log"

	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/admin/adminpb"
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial method connects to the service to the admin client
func ClientDial(cfg config.Config) (pb.AdminServiceClient, error) {
	grpc, err := grpc.Dial(cfg.ADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to user client at port : %s", cfg.USERPORT)
	return pb.NewAdminServiceClient(grpc), nil
}
