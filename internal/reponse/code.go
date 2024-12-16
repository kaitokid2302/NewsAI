package reponse

const (

	// user

	GetUserSuccess = 001
	GetUserFail    = 002
	// RegisterSucess auth
	RegisterSucess = 101
	RegisterFail   = 102

	// LoginSucess login
	LoginSucess = 201
	LoginFail   = 202

	// OTPVerifySucess otp
	OTPVerifySucess = 301
	OTPVerifyFail   = 302

	// ResendOTPSucess resend otp
	ResendOTPSucess = 401
	ResendOTPFail   = 402

	// JWTCreateFail jwt
	JWTCreateFail   = 501
	JWTCreateSucess = 502
	JWTVerifySucess = 503
	JWTVerifyFail   = 504

	// update

	UpdateUserSuccess = 601
	UpdateUserFail    = 602

	// topic

	SubscribeTopicSuccess = 701
	SubscribeTopicFail    = 702

	// Unsubscribe
	UnsubscribeTopicFail    = 801
	UnsubscribeTopicSuccess = 802

	// AllTopic
	AllTopicSuccess = 901
	AllTopicFail    = 902

	// Article
	GetArticleSuccess = 1001
	GetArticleFail    = 1002
)

var msg = map[int]string{
	GetUserSuccess:          "Get user information successfully",
	GetUserFail:             "Can not get user information",
	RegisterSucess:          "User has been registered successfully, OTP has been sent to your email",
	RegisterFail:            "Can not register user",
	LoginFail:               "Can not login",
	LoginSucess:             "Login successfully",
	OTPVerifySucess:         "OTP has been verified successfully",
	OTPVerifyFail:           "Can not verify OTP",
	ResendOTPSucess:         "OTP has been sent successfully again",
	ResendOTPFail:           "Can not send OTP again",
	JWTCreateFail:           "Can not create token",
	JWTCreateSucess:         "Token has been created successfully",
	JWTVerifySucess:         "Token has been verified successfully",
	JWTVerifyFail:           "Can not verify token",
	UpdateUserFail:          "Can not update user information",
	UpdateUserSuccess:       "Updated user information",
	SubscribeTopicSuccess:   "Subscribed topic successfully",
	SubscribeTopicFail:      "Can not subscribe topic",
	UnsubscribeTopicFail:    "Can not unsubscribe topic",
	UnsubscribeTopicSuccess: "Unsubscribed topic successfully",
	AllTopicSuccess:         "Get all topic successfully",
	AllTopicFail:            "Can not get all topic",
	GetArticleSuccess:       "Get article successfully",
	GetArticleFail:          "Can not get article",
}
