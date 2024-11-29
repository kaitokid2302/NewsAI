package auth

type OTPVerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   int    `json:"otp" binding:"required"`
}
