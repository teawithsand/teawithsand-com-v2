package user

type RegisterUserData struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	CaptchaResponse string `json:"captchaResponse"`
}

type RegisterRequest struct {
	RegisterUserData
}
type RegisterResponse struct {
	// Empty
}

type ConfirmRegistrationData struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type ConfirmRegistrationRequest struct {
	ConfirmRegistrationData
}
type ConfirmRegistrationResponse struct {
	// Empty
}

type InnerRegisterResponse struct {
	RegisterMailData RegisterMailData
}
