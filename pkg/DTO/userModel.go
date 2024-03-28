package dto

// User represents the user data structure.
type User struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Mobile       string `json:"mobile"`
	ReferralCode string `json:"referral_code"`
}

// OTP represents the structure for OTP verification.
type OTP struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

// Login represents the login credentials.
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Password represents the structure for updating password.
type Password struct {
	Old     string `json:"old_password"`
	New     string `json:"new_password"`
	Confirm string `json:"confirm_password"`
}
