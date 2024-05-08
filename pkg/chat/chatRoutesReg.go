package chat

import (
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/chat/handler"
	"github.com/gin-gonic/gin"
)

// Chat handles the websocket connection in the chat page.
func (c *Chat) Chat(ctx *gin.Context) {
	handler.HandleWebSocketConnection(ctx, c.client, c.userClient)
}

// ChatPage handles to load the chat page
func (c *Chat) ChatPage(ctx *gin.Context) {
	handler.ChatPage(ctx, c.client)
}
