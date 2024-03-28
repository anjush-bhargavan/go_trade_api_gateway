package admin

import (
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/admin/handler"
	"github.com/gin-gonic/gin"
)

// AdminLogin handles the admin login request.
func (a *Admin) AdminLogin(ctx *gin.Context) {
	handler.AdminLoginHandler(ctx, a.Client)
}

// AddCategory handles the admin to add category request.
func (a *Admin) AddCategory(ctx *gin.Context) {
	handler.AddCategoryHandler(ctx, a.Client)
}

// EditCategory handles the admin to edit category request.
func (a *Admin) EditCategory(ctx *gin.Context) {
	handler.EditCategoryHandler(ctx, a.Client)
}

// RemoveCategory handles the admin to remove category request.
func (a *Admin) RemoveCategory(ctx *gin.Context) {
	handler.RemoveCategoryHandler(ctx, a.Client)
}

// FindCategory handles the admin to find category request.
func (a *Admin) FindCategory(ctx *gin.Context) {
	handler.FindCategoryHandler(ctx, a.Client)
}

// FindCategories handles the admin to find all categories request.
func (a *Admin) FindCategories(ctx *gin.Context) {
	handler.FindAllCategoryHandler(ctx, a.Client)
}

// RemoveProduct handles the admin to remove product request.
func (a *Admin) RemoveProduct(ctx *gin.Context) {
	handler.RemoveProductHandler(ctx, a.Client)
}

// FindProduct handles the admin to find the product request.
func (a *Admin) FindProduct(ctx *gin.Context) {
	handler.FindProductHandler(ctx, a.Client)
}

// FindProducts handles the admin to find all products request.
func (a *Admin) FindProducts(ctx *gin.Context) {
	handler.FindAllProductHandler(ctx, a.Client)
}


// BlockUser handles the admin to block user request.
func (a *Admin) BlockUser(ctx *gin.Context) {
	handler.BlockUserHandler(ctx, a.Client)
}