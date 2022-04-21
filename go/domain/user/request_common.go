package user

type LoginUserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequest struct {
	LoginUserData
}
type LoginResponse struct {
	User  SecretUserProjection `json:"user"`
	Token string               `json:"token"`
}

type InvalidateTokensRequest struct {
	// Empty
}

type InvalidateTokensResponse struct {
	// Empty
}

type ChangePasswordData struct {
	Password string `json:"password"`
}
type ChangePasswordRequest struct {
	ChangePasswordData
}

type ChangePasswordResponse struct {
	// Empty
}

type InitResetPasswordRequest struct {
	Email           string `json:"email"`
	CaptchaResponse string `json:"captchaResponse"`
}

type InnerInitPasswordResetRequest struct {
	InitPasswordResetMailData InitPasswordResetMailData
}

type InitResetPasswordResponse struct {
	// Empty
}

type ResetPasswordData struct {
	ChangePasswordData

	Email string `json:"email"`
	Token string `json:"token"`
}

type ResetPasswordRequest struct {
	ResetPasswordData
}

type ResetPasswordResponse struct {
	// Empty
}

type ChangeEmailData struct {
	Email           string `json:"email"`
	CaptchaResponse string `json:"captchaResponse"`
}

type ChangeEmailRequest struct {
	ChangeEmailData
}

type ChangeEmailResponse struct {
	// Empty
}

type ConfirmEmailRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
type ConfirmEmailResponse struct {
	// Empty
}
