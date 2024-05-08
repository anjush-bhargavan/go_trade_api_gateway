package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	dto "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/DTO"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// // BeASellerHandler handles the user to register as a seller request.
// func BeASellerHandler(c *gin.Context, client pb.UserServiceClient) {
// 	timeout := time.Second * 1000
// 	ctx, cancel := context.WithTimeout(c, timeout)
// 	defer cancel()

// 	id, ok := c.Get("user_id")
// 	if !ok {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
// 			"Message": "error while user id from context",
// 			"Error":   ""})
// 		return
// 	}

// 	userID, ok := id.(uint)
// 	if !ok {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
// 			"Message": "error while user id converting",
// 			"Error":   ""})
// 		return
// 	}

// 	response, err := client.BeASeller(ctx, &pb.ID{
// 		ID: uint32(userID),
// 	})

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
// 			"Message": "error in client response",
// 			"Data":    response,
// 			"Error":   err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusAccepted, gin.H{
// 		"Status":  http.StatusAccepted,
// 		"Message": "seller created successfully",
// 		"Data":    response,
// 	})
// }

// AddProductHandler handles the seller to add  product request.
func AddProductHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	var product dto.Product

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	response, err := client.AddProduct(ctx, &pb.UserProduct{
		Seller_ID: uint32(userID),
		Name:      product.Name,
		Category: &pb.UserCategory{
			Category_ID: uint32(product.Category),
		},
		Base_Price:         uint32(product.BasePrice),
		Description:        product.Description,
		Bidding_Type:       product.BiddingType,
		Bidding_Start_Time: product.BiddingStartTime,
		Bidding_End_Time:   product.BiddingEndTime,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "product created successfully",
		"Data":    response,
	})
}

// EditProductHandler handles the seller to edit  product request.
func EditProductHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	ProductIDString := c.Param("id")
	ProductID, err := strconv.Atoi(ProductIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	var product dto.Product

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	response, err := client.EditProductUser(ctx, &pb.UserProduct{
		Product_ID: uint32(ProductID),
		Seller_ID: uint32(userID),
		Name:      product.Name,
		Category: &pb.UserCategory{
			Category_ID: uint32(product.Category),
		},
		Base_Price:         uint32(product.BasePrice),
		Description:        product.Description,
		Bidding_Type:       product.BiddingType,
		Bidding_Start_Time: product.BiddingStartTime,
		Bidding_End_Time:   product.BiddingEndTime,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "product edited successfully",
		"Data":    response,
	})
}

// RemoveProductHandler handles the seller to remove  product request.
func RemoveProductHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	ProductIDString := c.Param("id")
	ProductID, err := strconv.Atoi(ProductIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.RemoveProductUser(ctx, &pb.IDs{
		ID:      uint32(ProductID),
		User_ID: uint32(userID),
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
