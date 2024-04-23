package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/config"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// UserPaymentHandler handles the user to pay the product request.
func UserPaymentHandler(c *gin.Context, client pb.UserServiceClient, cnfg *config.Config) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	userIDString := c.Query("id")
	productIDString := c.Query("product_id")

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting userID to int",
			"Error":   err.Error()})
		return
	}
	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting productID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.UserCreatePayment(ctx, &pb.UserBid{
		User_ID:    uint32(userID),
		Product_ID: uint32(productID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}
	fmt.Println(cnfg.RAZORPAY)
	c.HTML(http.StatusOK, "app.html", gin.H{
		"userID":     userID,
		"totalPrice": response.Amount,
		"productID":  response.Product_ID,
		"orderID":    response.Order_ID,
		"email":      response.User_Email,
		"userName":   response.User_Name,
		"rpay":       cnfg.RAZORPAY,
	})
}

// UserPaymentSuccessHandler handles the user to payment success request.
func UserPaymentSuccessHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	userIDString := c.Query("user_id")
	// orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	productIDString := c.Query("product_id")
	// signature := c.Query("signature")
	paymentAmount := c.Query("total")

	amount, err := strconv.ParseFloat(paymentAmount, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting amount to float",
			"Error":   err.Error()})
		return
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting userID to int",
			"Error":   err.Error()})
		return
	}
	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in converting productID to int",
			"Error":   err.Error()})
		return
	}
	fmt.Println(productID,productIDString)

	_, err = client.UserPaymentSuccess(ctx, &pb.UserPayment{
		User_ID:    uint32(userID),
		Product_ID: uint32(productID),
		Payment_ID: paymentID,
		Amount:     float32(amount),
	})
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true})
}

// SuccessPage renders the success page.
func SuccessPageHandler(c *gin.Context) {
	pID := c.Query("id")
	userID := c.Query("user_id")
	// fmt.Println(pID)
	// fmt.Println("Fully successful")

	c.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": pID,
		"userID":    userID,
	})
}
