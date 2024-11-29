package auth

type OTPVerificationRequest struct {
	Email string `json:"email" binding:"required,email" form:"email"`
	OTP   int    `json:"otp" binding:"required" form:"otp"`
}
