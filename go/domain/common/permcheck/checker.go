package permcheck

import "github.com/teawithsand/webpage/util/voter"

func MakeChecker() (checker voter.Checker, err error) {
	var voters []voter.Voter
	voters = registerUserVoters(voters)
	voters = registerLangkaVoters(voters)

	checker = voter.NewSimpleChecker(voters)
	return
}
