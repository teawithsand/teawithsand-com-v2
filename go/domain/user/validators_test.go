package user_test

import (
	"context"
	"testing"

	"github.com/teawithsand/webpage/domain/user"
)

func TestValidationsOK(t *testing.T) {
	validators := user.DefaultValidators
	ctx := context.Background()
	t.Run("emailChangeData", func(t *testing.T) {
		err := validators.ChangeEmailDataValidator.Validate(ctx, user.ChangeEmailData{
			Email: "validmail@gmail.com",
		})
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("registerUserData", func(t *testing.T) {
		err := validators.RegistrationCreateDataValidator.Validate(ctx, user.RegisterUserData{
			Username:        "validusername",
			Email:           "validemail@gmail.com",
			CaptchaResponse: "somecaptcharesponse",
		})
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("confirmRegistrationData", func(t *testing.T) {
		err := validators.ConfirmRegistrationValidator.Validate(ctx, user.ConfirmRegistrationData{
			Token:    "sometoken",
			Password: "somevalidpassword",
		})
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("changeEmailData", func(t *testing.T) {
		err := validators.ChangeEmailDataValidator.Validate(ctx, user.ChangeEmailData{
			Email: "validmail@gmail.com",
		})
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("changePasswordData", func(t *testing.T) {
		err := validators.ChangePasswordDataValidator.Validate(ctx, user.ChangePasswordData{
			Password: "asdfasdfasdf1234!!!AA",
		})
		if err != nil {
			t.Error(err)
			return
		}
	})
}
