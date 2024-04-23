package handler


import (
	"context"
	"net/http"
	"strconv"
	"time"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/admin/adminpb"
	"github.com/gin-gonic/gin"
)

// OrderHistoryHandler function will find all orders request to client.
func OrderHistoryHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()


	response, err := client.FetchOrders(ctx, &pb.AdminNoParam{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "orders fetched successfully",
		"Data":    response,
	})
}

// UserOrdersHandler function will find all orders by user request to client.
func UserOrdersHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id := c.Param("id")
	userID ,err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting id to int",
			"Error":   err.Error()})
		return
	}
	response, err := client.FetchOrderByUser(ctx, &pb.AdID{ID: uint32(userID)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "orders fetched successfully",
		"Data":    response,
	})
}
// OrderHandler function will send find the order request to client.
func OrderHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id := c.Param("id")
	orderID ,err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting id to int",
			"Error":   err.Error()})
		return
	}
	response, err := client.FetchOrder(ctx, &pb.AdID{ID: uint32(orderID)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "orders fetched successfully",
		"Data":    response,
	})
}