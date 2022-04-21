package voter

import (
	"context"
	"errors"
)

// NewSimpleChecker, which requires at least one voter to vote for to grant access.
// Denies access if no voters given voting it.
func NewSimpleChecker(voters []Voter) Checker {
	return &simpleChecker{
		voters: voters,
	}
}

type simpleChecker struct {
	voters []Voter
}

func (checker *simpleChecker) Check(ctx context.Context, voting Voting) (err error) {
	for _, v := range checker.voters {
		var ok bool
		ok, err = v.Supports(ctx, voting)
		if err != nil {
			return
		}

		if !ok {
			continue
		}

		err = v.Vote(ctx, voting)
		if errors.Is(err, ErrAccessDenied) {
			continue
		} else {
			// return any other error
			return
		}
	}

	err = ErrAccessDenied
	return
}
