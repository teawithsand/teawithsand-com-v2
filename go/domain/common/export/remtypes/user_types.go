package remtypes

import (
	"github.com/teawithsand/webpage/domain/user"
)

func registerUserTypes() {
	addType(user.RegisterUserData{})
	addType(user.RegisterRequest{})
	addType(user.RegisterResponse{})

	addType(user.ConfirmRegistrationData{})
	addType(user.ConfirmRegistrationRequest{})
	addType(user.ConfirmRegistrationResponse{})

	addType(user.LoginUserData{})
	addType(user.LoginRequest{})
	addType(user.LoginResponse{})

	addType(user.InvalidateTokensRequest{})
	addType(user.InvalidateTokensResponse{})

	addType(user.ChangePasswordRequest{})
	addType(user.ChangePasswordResponse{})

	addType(user.InitResetPasswordRequest{})
	addType(user.InitResetPasswordResponse{})

	addType(user.ResetPasswordData{})
	addType(user.ResetPasswordRequest{})
	addType(user.ResetPasswordResponse{})

	addType(user.ChangeEmailData{})
	addType(user.ChangeEmailRequest{})
	addType(user.ChangeEmailResponse{})

	addType(user.ConfirmEmailRequest{})
	addType(user.ConfirmEmailResponse{})

	addType(user.GetUserSecretProfileRequest{})
	addType(user.GetUserSecretProfileResponse{})

	addType(user.PublicUserProjection{})
	addType(user.SecretUserProjection{})
	addType(user.UserReferenceProjection{})

	addVar("AuthHeaderName", user.AuthHeaderName)
	addVar("RemoteUserConfirmRegistrationEndpoint", UserConfirmRegistrationEndpoint)
}
