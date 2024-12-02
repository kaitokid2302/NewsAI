package reponse

const (
	// auth
	RegisterSucess = 101
	RegisterFail   = 102

	// login
	LoginSucess = 201
	LoginFail   = 202

	// otp
	OTPVerifySucess = 301
	OTPVerifyFail   = 302

	// resend otp
	ResendOTPSucess = 401
	ResendOTPFail   = 402

	// jwt
	JWTCreateFail   = 501
	JWTCreateSucess = 502
	JWTVerifySucess = 503
	JWTVerifyFail   = 504
)

var msg = map[int]string{
	RegisterSucess:  "User has been registered successfully",
	RegisterFail:    "Can not register user",
	LoginFail:       "Can not login",
	LoginSucess:     "Login successfully",
	OTPVerifySucess: "OTP has been verified successfully",
	OTPVerifyFail:   "Can not verify OTP",
	ResendOTPSucess: "OTP has been sent successfully again",
	ResendOTPFail:   "Can not send OTP again",
	JWTCreateFail:   "Can not create token",
	JWTCreateSucess: "Token has been created successfully",
	JWTVerifySucess: "Token has been verified successfully",
	JWTVerifyFail:   "Can not verify token",
}
