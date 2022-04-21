package permcheck

import (
	"context"

	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/voter"
)

func registerUserVoters(voters []voter.Voter) (res []voter.Voter) {
	voters = append(voters, &voter.OperationsVoter{
		Voter: voter.VoterFunc(func(ctx context.Context, voting voter.Voting) (err error) {
			if voting.Object == nil {
				return
			}

			u := voting.Object.(*user.User)
			if u == nil {
				return
			}
			if !u.Lifecycle.DeletedAt.IsZero() || !u.Lifecycle.LockedAt.IsZero() {
				err = voter.ErrAccessDenied
				return
			}

			return
		}),
		Operations: map[voter.Operation]struct{}{
			user.VoteUseApp: {},
		},
	})

	voters = append(voters, &voter.OperationsVoter{
		Voter: voter.VoterFunc(func(ctx context.Context, voting voter.Voting) (err error) {
			return
		}),
		Operations: map[voter.Operation]struct{}{
			user.VoteRegister:            {},
			user.VoteLogin:               {},
			user.VoteConfirmRegistration: {},
			user.VoteConfirmEmail:        {},
			user.VoteInitResetPassword:   {},
			user.VoteResetPassword:       {},
		},
	})

	voters = append(voters, &voter.OperationsVoter{
		Voter: voter.VoterFunc(func(ctx context.Context, voting voter.Voting) (err error) {
			if voting.Object == nil {
				err = voter.ErrAccessDenied
				return
			}

			u := voting.Object.(*user.User)
			if u == nil {
				err = voter.ErrAccessDenied
				return
			}
			current := user.GetUser(ctx)
			if current == nil {
				err = voter.ErrAccessDenied
				return
			}

			if u.ID != current.ID {
				err = voter.ErrAccessDenied
				return
			}

			return
		}),
		Operations: map[voter.Operation]struct{}{
			user.VoteChangeEmail:       {},
			user.VoteChangePassword:    {},
			user.VoteInvalidateTokens:  {},
			user.VoteShowSecretProfile: {},
		},
	})

	res = voters
	return
}
