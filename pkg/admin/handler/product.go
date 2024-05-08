package handler

import (
	"context"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/admin/adminpb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// RemoveProductHandler function will send remove Product request to client.
func RemoveProductHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	ProductIDString := c.Param("id")
	ProductID, err := strconv.Atoi(ProductIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.RemoveProduct(ctx, &pb.AdID{
		ID: uint32(ProductID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Product removed successfully",
		"Data":    response,
	})
}

// FindProductHandler function will send find Product request to client.
func FindProductHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	ProductIDString := c.Param("id")
	ProductID, err := strconv.Atoi(ProductIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.FindProductByID(ctx, &pb.AdID{
		ID: uint32(ProductID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Product fetched successfully",
		"Data":    response,
	})
}

// FindAllProductHandler function will send find all Product request to client.
func FindAllProductHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.FindAllProducts(ctx, &pb.AdminNoParam{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Products fetched successfully",
		"Data":    response,
	})
}
