package user

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial method connects to the service to the user client
func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	grpc, err := grpc.Dial(cfg.USERPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to user client at port : %s", cfg.USERPORT)
	return pb.NewUserServiceClient(grpc), nil
}
