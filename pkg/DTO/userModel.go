package dto

// User represents the user data structure.
type User struct {
	UserName     string `json:"user_name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Mobile       string `json:"mobile" validate:"required"`
	ReferralCode string `json:"referral_code"`
}

// OTP represents the structure for OTP verification.
type OTP struct {
	Email string `json:"email" validate:"required"`
	OTP   string `json:"otp" validate:"required"`
}

// Login represents the login credentials.
type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Password represents the structure for updating password.
type Password struct {
	Old     string `json:"old_password" validate:"required"`
	New     string `json:"new_password" validate:"required"`
	Confirm string `json:"confirm_password" validate:"required"`
}
