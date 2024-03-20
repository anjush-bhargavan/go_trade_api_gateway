package user

import (
	"github.com/anjush-bhargavan/go_trade_api_gateway/pkg/user/handler"
	"github.com/gin-gonic/gin"
)

func (u *User) UserSignup(ctx *gin.Context) {
	handler.UserSignupHandler(ctx, u.client)
}

func (u *User) UserVerify(ctx *gin.Context) {
	handler.VerificationHandler(ctx, u.client)
}

func (u *User) UserLogin(ctx *gin.Context) {
	handler.UserLoginHandler(ctx, u.client)
}

func (u *User) AddAddress(ctx *gin.Context) {
	handler.AddAddressHandler(ctx, u.client)
}

func (u *User) EditAddress(ctx *gin.Context) {
	handler.EditAddressHandler(ctx, u.client)
}

func (u *User) RemoveAddress(ctx *gin.Context) {
	handler.RemoveAddressHandler(ctx, u.client)
}

func (u *User) ViewProfile(ctx *gin.Context) {
	handler.ViewProfileHandler(ctx, u.client)
}

func (u *User) EditProfile(ctx *gin.Context) {
	handler.EditProfileHandler(ctx, u.client)
}

func (u *User) ChangePassword(ctx *gin.Context) {
	handler.ChangePasswordHandler(ctx, u.client)
}

func (u *User) BeASeller(ctx *gin.Context) {
	handler.BeASellerHandler(ctx, u.client)
}

func (u *User) AddProduct(ctx *gin.Context) {
	handler.AddProductHandler(ctx, u.client)
}
