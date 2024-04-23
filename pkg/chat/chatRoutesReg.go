package chat

import (
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/chat/handler"
	"github.com/gin-gonic/gin"
)

// AddBid handles the request to add a bid to the product by the user.
func (c *Chat) Chat(ctx *gin.Context) {
	handler.HandleWebSocketConnection(ctx, c.client,c.userClient)
}

// AddBid handles the request to add a bid to the product by the user.
func (c *Chat) ChatPage(ctx *gin.Context) {
	handler.ChatPage(ctx,c.client)
}
