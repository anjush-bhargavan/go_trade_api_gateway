package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	dto "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/DTO"
	pb "github.com/anjush-bhargavan/go_trade_api_gateway/pkg/admin/adminpb"
	"github.com/gin-gonic/gin"
)

// AddCategoryHandler function will send add category request to client.
func AddCategoryHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var category dto.Category
	if err := c.BindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	response, err := client.AddCategory(ctx, &pb.AdminCategory{
		Name: category.Name,
		Description: category.Description,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "category added successfully",
		"Data":    response,
	})
}

// EditCategoryHandler function will send edit category request to client.
func EditCategoryHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var category dto.Category
	if err := c.BindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	categoryIDString := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.EditCategoryAdmin(ctx, &pb.AdminCategory{
		Category_ID: uint32(categoryID),
		Name: category.Name,
		Description: category.Description,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "category edited successfully",
		"Data":    response,
	})
}

// RemoveCategoryHandler function will send remove category request to client.
func RemoveCategoryHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	categoryIDString := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.RemoveCategoryAdmin(ctx, &pb.AdID{
		ID: uint32(categoryID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "category removed successfully",
		"Data":    response,
	})
}

// FindCategoryHandler function will send find category request to client.
func FindCategoryHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	categoryIDString := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.FindCategory(ctx, &pb.AdID{
		ID: uint32(categoryID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "category fetched successfully",
		"Data":    response,
	})
}

// FindAllCategoryHandler function will send find all category request to client.
func FindAllCategoryHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.FindCategories(ctx, &pb.AdminNoParam{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "category fetched successfully",
		"Data":    response,
	})
}