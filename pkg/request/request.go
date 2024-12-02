package request

type OTPVerificationRequest struct {
	Email string `json:"email" binding:"required,email" form:"email"`
	OTP   int    `json:"otp" binding:"required" form:"otp"`
}

type ResendOTPRequest struct {
	Email string `json:"email" binding:"required,email" form:"email"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" form:"name"`
	Email    string `json:"email" binding:"required,email" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}
