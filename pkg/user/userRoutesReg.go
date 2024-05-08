package user

import (
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/handler"
	"github.com/gin-gonic/gin"
)

// UserSignup handles the user signup request.
func (u *User) UserSignup(ctx *gin.Context) {
	handler.UserSignupHandler(ctx, u.client)
}

// UserVerify handles the user verification request.
func (u *User) UserVerify(ctx *gin.Context) {
	handler.VerificationHandler(ctx, u.client)
}

// UserLogin handles the user login request.
func (u *User) UserLogin(ctx *gin.Context) {
	handler.UserLoginHandler(ctx, u.client)
}

// AddAddress handles the request to add a new address for the user.
func (u *User) AddAddress(ctx *gin.Context) {
	handler.AddAddressHandler(ctx, u.client)
}

// EditAddress handles the request to edit an existing address for the user.
func (u *User) EditAddress(ctx *gin.Context) {
	handler.EditAddressHandler(ctx, u.client)
}

// RemoveAddress handles the request to remove an existing address for the user.
func (u *User) RemoveAddress(ctx *gin.Context) {
	handler.RemoveAddressHandler(ctx, u.client)
}

// ViewProfile handles the request to view the user's profile.
func (u *User) ViewProfile(ctx *gin.Context) {
	handler.ViewProfileHandler(ctx, u.client)
}

// EditProfile handles the request to edit the user's profile.
func (u *User) EditProfile(ctx *gin.Context) {
	handler.EditProfileHandler(ctx, u.client)
}

// ChangePassword handles the request to change the user's password.
func (u *User) ChangePassword(ctx *gin.Context) {
	handler.ChangePasswordHandler(ctx, u.client)
}

// // BeASeller handles the request for the user to become a seller.
// func (u *User) BeASeller(ctx *gin.Context) {
// 	handler.BeASellerHandler(ctx, u.client)
// }

// AddProduct handles the request to add a new product by the user.
func (u *User) AddProduct(ctx *gin.Context) {
	handler.AddProductHandler(ctx, u.client)
}

// EditProduct handles the request to edit the product by the user.
func (u *User) EditProduct(ctx *gin.Context) {
	handler.EditProductHandler(ctx, u.client)
}

// RemoveProduct handles the request to add a new product by the user.
func (u *User) RemoveProduct(ctx *gin.Context) {
	handler.RemoveProductHandler(ctx, u.client)
}

// FindProduct handles the request to find product by the user.
func (u *User) FindProduct(ctx *gin.Context) {
	handler.UserFindProductHandler(ctx, u.client)
}

// FindAllProducts handles the request to find all products by the user.
func (u *User) FindAllProducts(ctx *gin.Context) {
	handler.UserFindAllProductHandler(ctx, u.client)
}

// FindProductsByCategory handles the request to find products by the category id for user.
func (u *User) FindProductsByCategory(ctx *gin.Context) {
	handler.UserFindAllProductByCategoryHandler(ctx, u.client)
}

// FindCategory handles the request to find category by the user.
func (u *User) FindCategory(ctx *gin.Context) {
	handler.UserFindCategoryHandler(ctx, u.client)
}

// FindAllCategories handles the request to find all categories by the user.
func (u *User) FindAllCategories(ctx *gin.Context) {
	handler.UserFindAllCategoryHandler(ctx, u.client)
}

// AddToWatchlist handles the request to add a category to the users watchlist by the user.
func (u *User) AddToWatchlist(ctx *gin.Context) {
	handler.AddTOWatchlistHandler(ctx, u.client)
}

// ViewWatchlist handles the request the users to view  watchlist by the user.
func (u *User) ViewWatchlist(ctx *gin.Context) {
	handler.ViewWatchlistHandler(ctx, u.client)
}

// AddBid handles the request to add a bid to the product by the user.
func (u *User) AddBid(ctx *gin.Context) {
	handler.AddBidHandler(ctx, u.client)
}

// GetBids handles the request the users to view  the bids on the product by the user.
func (u *User) GetBids(ctx *gin.Context) {
	handler.GetBidsHandler(ctx, u.client)
}

// Payment to load payment page and do payment using razorpay.
func (u *User) Payment(ctx *gin.Context) {
	handler.UserPaymentHandler(ctx, u.client, u.cfg)
}

// PaymentSuccess for completeing the procedure after payment success
func (u *User) PaymentSuccess(ctx *gin.Context) {
	handler.UserPaymentSuccessHandler(ctx, u.client)
}

// SuccessPage is to load the payment completed page.
func (u *User) SuccessPage(ctx *gin.Context) {
	handler.SuccessPageHandler(ctx)
}
