package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/captcha"
	"github.com/teawithsand/webpage/util/cext"
	"github.com/teawithsand/webpage/util/dbutil"
	"github.com/teawithsand/webpage/util/mailer"
	"github.com/teawithsand/webpage/util/phash"
	"github.com/teawithsand/webpage/util/token"
	"github.com/teawithsand/webpage/util/voter"
)

const RefreshTokenLiveness = time.Hour * 24 * 365

type UserService struct {
	Checker voter.Checker

	UserRepository         UserRepository
	RegistrationRepository RegistrationRepository
	PasswordHasher         phash.Hasher

	RefreshTokenNonceGenerator  util.NonceGenerator
	RegistrationTokenGenerator  util.NonceGenerator
	PasswordResetTokenGenerator util.NonceGenerator
	EmailConfirmTokenGenerator  util.NonceGenerator

	InitPasswordResetMailer mailer.Mailer
	RegistrationMailer      mailer.Mailer

	RefreshTokenManager token.Manager

	CaptchaValidator captcha.Validator

	// For sake of simplicity,do not use separate DI injects for validators
	Validators Validators
}

func (svc *UserService) InnerRegisterUser(ctx context.Context, req RegisterRequest) (res InnerRegisterResponse, err error) {
	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteRegister,
		Object:    req.RegisterUserData,
	})
	if err != nil {
		return
	}

	err = svc.CaptchaValidator.ValidateCaptchaResponse(ctx, req.CaptchaResponse)
	if err != nil {
		return
	}

	err = svc.Validators.RegistrationCreateDataValidator.Validate(ctx, req.RegisterUserData)
	if err != nil {
		return
	}

	token, err := svc.RegistrationTokenGenerator.GenerateNonce(ctx)
	if err != nil {
		return
	}

	ok, err := svc.UserRepository.FindOne(ctx, NewPublicNameUserQuery(req.Username), nil)
	if err != nil {
		return
	}
	if ok {
		err = ErrLoginInUse
		return
	}

	ok, err = svc.UserRepository.FindOne(ctx, NewEmailUserQuery(req.Email), nil)
	if err != nil {
		return
	}
	if ok {
		err = ErrEmailInUse
		return
	}

	reg := Registration{
		ID:         db.NewID(),
		PublicName: req.Username,
		Email:      req.Email,
		Token:      token,
		CreatedAt:  cext.Now(ctx),
	}

	err = svc.RegistrationRepository.Create(ctx, reg)

	if db.IsUniqueViolatedError(err) {
		err = ErrLoginInUse
		return
	} else if err != nil {
		return
	}

	res.RegisterMailData = RegisterMailData{
		Username:  reg.PublicName,
		Token:     reg.Token,
		Email:     reg.Email,
		CreatedAt: reg.CreatedAt,
	}

	return
}

func (svc *UserService) RegisterUser(ctx context.Context, req RegisterRequest) (res RegisterResponse, err error) {
	ires, err := svc.InnerRegisterUser(ctx, req)
	if err != nil {
		return
	}
	err = svc.RegistrationMailer.SendEmail(ctx, ires.RegisterMailData)

	if err != nil {
		return
	}

	return
}

func (svc *UserService) ConfirmRegistration(ctx context.Context, req ConfirmRegistrationRequest) (res ConfirmRegistrationResponse, err error) {
	var reg Registration
	ok, err := svc.RegistrationRepository.FindOne(ctx, NewTokenRegistrationQuery(req.Token), &reg)
	if err != nil {
		return
	}
	if !ok {
		err = ErrRegistrationForTokenNotFound
		return
	}

	err = svc.Validators.ConfirmRegistrationValidator.Validate(ctx, req.ConfirmRegistrationData)
	if err != nil {
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteConfirmRegistration,
		Object:    &reg,
	})
	if err != nil {
		return
	}

	phash, err := svc.PasswordHasher.HashPassword(ctx, req.Password)
	if err != nil {
		return
	}

	user, err := reg.MakeUser(
		ctx,
		svc.RefreshTokenNonceGenerator,
		svc.PasswordResetTokenGenerator,
		phash,
	)
	if err != nil {
		return
	}

	err = svc.UserRepository.Create(ctx, *user)
	if db.IsUniqueViolatedError(err) {
		err = ErrUserAlreadyRegistered
		return
	} else if err != nil {
		return
	}

	err = svc.RegistrationRepository.DeleteMany(ctx, NewEmailRegistrationQuery(reg.Email))
	if err != nil {
		return
	}

	return
}

func (svc *UserService) LoginUser(ctx context.Context, req LoginRequest) (res LoginResponse, err error) {
	var user User
	ok, err := svc.UserRepository.FindOne(ctx, NewPublicNameUserQuery(req.Login), &user)
	if err != nil {
		return
	}
	if !ok {
		err = ErrUserForLoginNotFound
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteLogin,
		Object:    user,
	})
	if err != nil {
		return
	}

	err = svc.PasswordHasher.CheckPassword(ctx, req.Password, user.NativeAuthData.PasswordHash)
	if err != nil {
		err = ErrUserForLoginNotFound
		return
	}

	token, err := svc.RefreshTokenManager.IssueToken(ctx, RefreshToken{
		JWTStandardClaims: jwt.StandardClaims{
			ExpiresAt: cext.Now(ctx).Add(RefreshTokenLiveness).Unix(),
		},
		Nonce:  user.TokenAuthData.Nonce,
		UserID: user.ID,
	})
	if err != nil {
		return
	}

	res.User.Load(&user)

	res.Token = token
	return
}

func (svc *UserService) InvalidateTokens(ctx context.Context, req InvalidateTokensRequest) (res InvalidateTokensResponse, err error) {
	user := GetUser(ctx)

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteInvalidateTokens,
		Object:    user,
	})
	if err != nil {
		return
	}

	if user == nil {
		panic("user can't be nil here")
	}

	newNonce, err := svc.RefreshTokenNonceGenerator.GenerateNonce(ctx)
	if err != nil {
		return
	}

	err = svc.UserRepository.UpdateOne(ctx, dbutil.NewIDQuery(user.ID), NewTokenNonceUserUpdate(newNonce))
	if err != nil {
		return
	}

	return
}

func (svc *UserService) makeChangePasswordData(ctx context.Context, password string) (data interface{}, err error) {
	phash, err := svc.PasswordHasher.HashPassword(ctx, password)
	if err != nil {
		return
	}

	newTokenNonce, err := svc.RefreshTokenNonceGenerator.GenerateNonce(ctx)
	if err != nil {
		return
	}

	newPasswordResetToken, err := svc.PasswordResetTokenGenerator.GenerateNonce(ctx)
	if err != nil {
		return
	}

	data = NewPasswordAndTokensAndNoncesUserUpdate(
		phash,
		newTokenNonce,
		newPasswordResetToken,
	)
	return
}

func (svc *UserService) InnerInitPasswordReset(ctx context.Context, req InitResetPasswordRequest) (res InnerInitPasswordResetRequest, err error) {
	err = svc.CaptchaValidator.ValidateCaptchaResponse(ctx, req.CaptchaResponse)
	if err != nil {
		return
	}

	var user User
	ok, err := svc.UserRepository.FindOne(ctx, NewEmailUserQuery(req.Email), &user)
	if err != nil {
		return
	}

	if !ok {
		err = ErrNoUserForPasswordInit
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteInitResetPassword,
		Object:    user,
	})
	if err != nil {
		return
	}

	res.InitPasswordResetMailData = InitPasswordResetMailData{
		PublicName: user.PublicName,
		Token:      user.NativeAuthData.PasswordResetToken,
		Email:      user.Email.Email,
		CreatedAt:  cext.Now(ctx),
	}
	return
}

func (svc *UserService) InitPasswordReset(ctx context.Context, req InitResetPasswordRequest) (res InitResetPasswordResponse, err error) {
	ires, err := svc.InnerInitPasswordReset(ctx, req)
	if err != nil {
		return
	}

	err = svc.InitPasswordResetMailer.SendEmail(ctx, ires.InitPasswordResetMailData)
	if err != nil {
		return
	}

	return
}

func (svc *UserService) PasswordReset(ctx context.Context, req ResetPasswordRequest) (res ResetPasswordResponse, err error) {
	var user User
	ok, err := svc.UserRepository.FindOne(ctx, NewEmailUserQuery(req.Email), &user)
	if err != nil {
		return
	}
	if !ok {
		err = ErrPasswordResetFiled
		return
	}

	if !util.SafeStringEquals(user.NativeAuthData.PasswordResetToken, req.Token) {
		err = ErrPasswordResetFiled
		return
	}

	err = svc.Validators.ChangePasswordDataValidator.Validate(ctx, req.ResetPasswordData.ChangePasswordData)
	if err != nil {
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteResetPassword,
		Object:    &user,
	})
	if err != nil {
		return
	}

	updateData, err := svc.makeChangePasswordData(ctx, req.Password)
	if err != nil {
		return
	}

	err = svc.UserRepository.UpdateOne(
		ctx,
		dbutil.NewIDQuery(user.ID),
		updateData,
	)
	if err != nil {
		return
	}

	return
}

func (svc *UserService) ChangePassword(ctx context.Context, req ChangePasswordRequest) (res ChangePasswordResponse, err error) {
	user := GetUser(ctx)

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteChangePassword,
		Object:    user,
	})
	if err != nil {
		return
	}

	if user == nil {
		panic("user can't be nil here")
	}

	err = svc.Validators.ChangePasswordDataValidator.Validate(ctx, req.ChangePasswordData)
	if err != nil {
		return
	}

	updateData, err := svc.makeChangePasswordData(ctx, req.Password)
	if err != nil {
		return
	}

	err = svc.UserRepository.UpdateOne(
		ctx,
		dbutil.NewIDQuery(user.ID),
		updateData,
	)
	if err != nil {
		return
	}

	return
}

func (svc *UserService) ChangeEmail(ctx context.Context, req ChangeEmailRequest) (res ChangeEmailResponse, err error) {
	user := GetUser(ctx)

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteChangeEmail,
		Object:    user,
	})
	if err != nil {
		return
	}

	if user == nil {
		panic("user can't be nil here")
	}

	err = svc.Validators.ChangeEmailDataValidator.Validate(ctx, req.ChangeEmailData)
	if err != nil {
		return
	}

	confirmEmailToken, err := svc.EmailConfirmTokenGenerator.GenerateNonce(ctx)
	if err != nil {
		return
	}

	err = svc.UserRepository.UpdateOne(
		ctx,
		dbutil.NewIDQuery(user.ID),
		NewEmailAndNonceUserUpdate(req.Email, confirmEmailToken),
	)
	if err != nil {
		return
	}

	return
}

func (svc *UserService) ConfirmEmail(ctx context.Context, req ConfirmEmailRequest) (res ConfirmEmailResponse, err error) {
	var user User
	ok, err := svc.UserRepository.FindOne(ctx, NewEmailUserQuery(req.Email), &user)
	if err != nil {
		return
	}

	if !ok {
		err = ErrEmailConfirmFiled
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteConfirmEmail,
		Object:    &user,
	})
	if err != nil {
		return
	}

	if !user.Email.EmailConfirmedAt.IsZero() {
		err = ErrEmailConfirmedAlready
		return
	}

	if !util.SafeStringEquals(user.Email.EmailConfirmToken, req.Token) {
		err = ErrEmailConfirmFiled
		return
	}

	err = svc.UserRepository.UpdateOne(
		ctx,
		dbutil.NewIDQuery(user.ID),
		NewEmailConfirmedUserUpdate(cext.Now(ctx)),
	)
	if err != nil {
		return
	}

	return
}

func (svc *UserService) GetUserSecretProfile(ctx context.Context, req GetUserSecretProfileRequest) (res GetUserSecretProfileResponse, err error) {
	var user User
	ok, err := svc.UserRepository.FindOne(ctx, dbutil.NewIDQuery(req.ID), &user)
	if err != nil {
		return
	}

	if !ok {
		err = ErrProfileNotFound
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteShowSecretProfile,
		Object:    &user,
	})
	if err != nil {
		return
	}

	res.SecretUserProjection.Load(&user)
	return
}
