package user

import "github.com/teawithsand/webpage/util/voter"

const VoteUseApp = voter.Operation("VOTE_USER_APP")

const VoteChangeEmail = voter.Operation("VOTE_USER_CHANGE_EMAIL")
const VoteConfirmEmail = voter.Operation("VOTE_USER_CONFIRM_EMAIL")

const VoteLogin = voter.Operation("VOTE_USER_LOGIN")

const VoteInvalidateTokens = voter.Operation("VOTE_USER_INVALIDATE_TOKENS")

const VoteChangePassword = voter.Operation("VOTE_USER_CHANGE_PASSWORD")
const VoteInitResetPassword = voter.Operation("VOTE_USER_INIT_RESET_PASSWORD")
const VoteResetPassword = voter.Operation("VOTE_USER_RESET_PASSWORD")

const VoteRegister = voter.Operation("VOTE_USER_REGISTER")
const VoteConfirmRegistration = voter.Operation("VOTE_USER_CONFIRM_REGISTRATION")

const VoteShowSecretProfile = voter.Operation("VOTE_SHOW_SECRET_PROFILE")
