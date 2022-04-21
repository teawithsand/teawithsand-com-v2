package user_test

import (
	"context"
	"testing"

	"github.com/teawithsand/webpage/domain/domtestutil"
	"github.com/teawithsand/webpage/domain/domtestutil/fixture"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/testutil"
)

func TestUserService_Login(t *testing.T) {
	testutil.SetTestENV()

	tdi := domtestutil.MustConstructTestDI()
	defer tdi.Close()

	ctx := testutil.Context(t)
	ctx, closer, err := tdi.DB.MakeSession(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer closer.Close()

	err = fixture.Apply(ctx, tdi.DB)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := tdi.UserService.LoginUser(ctx, user.LoginRequest{
		LoginUserData: user.LoginUserData{
			Login:    fixture.MockUserOne.PublicName,
			Password: fixture.MockUserOne.NativeAuthData.PasswordHash,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	_ = res
}

func TestUserService_Register(t *testing.T) {
	testutil.SetTestENV()

	tdi := domtestutil.MustConstructTestDI()
	defer tdi.Close()

	ctx := testutil.Context(t)
	ctx, closer, err := tdi.DB.MakeSession(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer closer.Close()

	err = tdi.DB.Clear(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = tdi.UserService.RegisterUser(ctx, user.RegisterRequest{
		RegisterUserData: user.RegisterUserData{
			Username:        "validlogin",
			Email:           "validemail@example.com",
			CaptchaResponse: "iamnotarobot",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUserService_RegisterAndConfirm(t *testing.T) {
	testutil.SetTestENV()

	tdi := domtestutil.MustConstructTestDI()
	defer tdi.Close()

	ctx := testutil.Context(t)
	ctx, closer, err := tdi.DB.MakeSession(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer closer.Close()

	res, err := tdi.UserService.InnerRegisterUser(ctx, user.RegisterRequest{
		RegisterUserData: user.RegisterUserData{
			Username:        "validlogin",
			Email:           "validemail@example.com",
			CaptchaResponse: "iamnotarobot",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = tdi.UserService.ConfirmRegistration(ctx, user.ConfirmRegistrationRequest{
		ConfirmRegistrationData: user.ConfirmRegistrationData{
			Token:    res.RegisterMailData.Token,
			Password: "Validpassword1234!",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	var innerUser *user.User
	ok, err := tdi.UserRepository.FindOne(ctx, user.NewPublicNameUserQuery("validlogin"), &innerUser)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("user is nil, but it should not be")
		return
	}
}

func TestUserService_ChangePassword(t *testing.T) {
	testutil.SetTestENV()

	tdi := domtestutil.MustConstructTestDI()
	defer tdi.Close()

	ctx := context.Background()
	ctx, closer, err := tdi.DB.MakeSession(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer closer.Close()

	err = fixture.Apply(ctx, tdi.DB)
	if err != nil {
		t.Error(err)
		return
	}

	// TODO(teawithsand): implement this test
}

func TestUserService_GetSecretProfile(t *testing.T) {
	testutil.SetTestENV()

	tdi := domtestutil.MustConstructTestDI()
	defer tdi.Close()

	ctx := context.Background()
	ctx, closer, err := tdi.DB.MakeSession(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer closer.Close()

	err = fixture.Apply(ctx, tdi.DB)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = tdi.UserService.GetUserSecretProfile(ctx, user.GetUserSecretProfileRequest{
		ID: fixture.MockUserOne.ID,
	})
	if err == nil {
		t.Error("expected access denied here")
		return
	}

	ctx = user.PutUser(ctx, &fixture.MockUserOne)

	res, err := tdi.UserService.GetUserSecretProfile(ctx, user.GetUserSecretProfileRequest{
		ID: fixture.MockUserOne.ID,
	})
	if err != nil {
		t.Error(err)
		return
	}

	if res.SecretUserProjection == (user.SecretUserProjection{}) {
		t.Error("result must not be empty")
		return
	}

}
