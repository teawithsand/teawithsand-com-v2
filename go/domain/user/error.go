package user

import "errors"

// Register

var ErrLoginInUse = errors.New("domain/user: specified login is in use")
var ErrEmailInUse = errors.New("domain/user: specified email is already in use") // + ChangeEmail

// FinalizeRegistration

var ErrRegistrationForTokenNotFound = errors.New("domain/user: registration for token wasn't found. It may have expired")
var ErrUserAlreadyRegistered = errors.New("domain/user: user for given email was already registered and confirmed") // Confirm email

// Login

var ErrUserForLoginNotFound = errors.New("domain/user: user with given login and password was not found")

// Auth MW

var ErrInvalidAuthHeader = errors.New("domain/user: got invalid header with token")
var ErrTokenUserNotFound = errors.New("domain/user: user with id specified on token does not exist")
var ErrTokenUserNonceMismatch = errors.New("domain/user: given token has invalid nonce and can't be used")

// InitPasswordReset

var ErrNoUserForPasswordInit = errors.New("domain/user: user for specified email for initializing password reset does not exist")

// PasswordReset

var ErrPasswordResetFiled = errors.New("domain/user: password reset filed because token or email invalid")

// ConfirmEmail

var ErrEmailConfirmedAlready = errors.New("domain/user: email was already confirmed")
var ErrEmailConfirmFiled = errors.New("domain/user: email confirm filed due to invalid token or invalid email")

var ErrProfileNotFound = errors.New("domain/user: user profile was not found")
