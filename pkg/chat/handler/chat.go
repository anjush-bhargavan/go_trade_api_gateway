package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	dto "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/DTO"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/chat/pb"
	userpb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/userpb"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocketConnection(c *gin.Context, client pb.ChatServiceClient, userClient userpb.UserServiceClient) {

	ctx := c.Request.Context() // Get the request context

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		select {
		case <-ctx.Done():
			// Context canceled, stop processing messages
			log.Println("WebSocket connection closed")
			return
		default:
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			var message dto.Message
			err = json.Unmarshal(msg, &message)
			if err != nil {
				log.Println("Error decoding JSON:", err)
				continue
			}

			//checking the user and receiver IDs
			_, err = userClient.ViewProfile(ctx, &userpb.ID{ID: uint32(message.UserID)})
			if err != nil {
				log.Fatalf("error fetching user or invalid userID")
			}

			_, err = userClient.ViewProfile(ctx, &userpb.ID{ID: uint32(message.ReceiverID)})
			if err != nil {
				log.Fatalf("error fetching receiver or invalid receiverID")
			}
			// fmt.Println(user.User_Name, receiver.User_Name)
			// Process message (e.g., store it in a database)
			stream, err := client.Connect(ctx)
			if err != nil {
				log.Fatalf("error calling chat service")
			}
			ch := &clientHandle{
				stream:     stream,
				userID:     message.UserID,
				receiverID: message.ReceiverID,
			}
			// log.Printf("Received message from %d to %d: %s\n", message.UserID, message.ReceiverID, message.Message)

			// Echo message back to sender
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				break
			}

			go ch.sentMessage(message.Message)
			go ch.receiveMessage(conn, message.UserID, message.ReceiverID)
		}
	}
}

func ChatPage(c *gin.Context, client pb.ChatServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id := c.Query("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting id to int",
			"Error":   err.Error()})
		return
	}
	receiverId := c.Query("receiverId")
	receiverID, err := strconv.Atoi(receiverId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting id to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.FetchHistory(ctx, &pb.ChatID{
		User_ID:     uint32(userID),
		Receiver_ID: uint32(receiverID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in calling chat client",
			"Error":   err.Error()})
		return
	}
	c.HTML(http.StatusOK, "chat.html", gin.H{"response": response.Chats,"id":userID})
}

type clientHandle struct {
	userID     uint
	receiverID uint
	stream     pb.ChatService_ConnectClient
}

func (c *clientHandle) sentMessage(msg string) {
	message := &pb.Message{
		User_ID:     uint32(c.userID),
		Receiver_ID: uint32(c.receiverID),
		Content:     msg,
	}

	err := c.stream.Send(message)
	if err != nil {
		log.Printf("Error while sending message to server :: %v", err)
	}

}

// receive message
func (ch *clientHandle) receiveMessage(c *websocket.Conn, userID, receiverID uint) {

	//create a loop
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		if uint32(userID) == mssg.Receiver_ID && receiverID == uint(mssg.User_ID) {
			dom := &dto.Message{
				UserID:     uint(mssg.User_ID),
				ReceiverID: uint(mssg.Receiver_ID),
				Message:    mssg.Content,
			}
			msg, err := json.Marshal(dom)
			if err != nil {
				log.Println("Error decoding JSON:", err)
				continue
			}
			// err = c.WriteMessage(websocket.TextMessage, msg)
			// if err != nil {
			// 	log.Println("Error writing message:", err)
			// 	break
			// }
			// Write received message to WebSocket connection
			err = c.WriteMessage(websocket.TextMessage, (msg))
			if err != nil {
				log.Println("Error writing message:", err)
				break
			}
		}

		//print message to console
		// fmt.Printf("%s : %d to %d\n", mssg.Content, mssg.User_ID, mssg.Receiver_ID)

	}
}
