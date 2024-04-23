package chat

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/chat/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.ChatServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.CHATPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc chat client: %s, ", cfg.CHATPORT)
		return nil, err
	}
	
	log.Printf("Succesfully connected to chat client at port: %v", cfg.CHATPORT)
	return pb.NewChatServiceClient(grpc), nil
}