package permcheck

import (
	"context"

	"github.com/teawithsand/webpage/domain/langka"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/voter"
)

func registerLangkaVoters(voters []voter.Voter) (res []voter.Voter) {
	voters = append(voters, &voter.OperationsVoter{
		Voter: voter.VoterFunc(func(ctx context.Context, voting voter.Voting) (err error) {
			u := user.GetUser(ctx)
			if u == nil {
				err = voter.ErrAccessDenied
				return
			}

			ws := voting.Object.(*langka.WordSet)
			if ws == nil {
				err = voter.ErrAccessDenied
				return
			}

			if ws.Owner.ID != u.ID {
				err = voter.ErrAccessDenied
				return
			}
			return
		}),
		Operations: map[voter.Operation]struct{}{
			langka.VoteEditWordSet:      {},
			langka.VoteDeleteWordSet:    {},
			langka.VotePublishWordSet:   {},
			langka.VoteUnpublishWordSet: {},
		},
	})

	voters = append(voters, &voter.OperationsVoter{
		Voter: voter.VoterFunc(func(ctx context.Context, voting voter.Voting) (err error) {
			u := user.GetUser(ctx)
			if u == nil {
				err = voter.ErrAccessDenied
				return
			}
			return
		}),
		Operations: map[voter.Operation]struct{}{
			langka.VoteCreateWordSet:       {},
			langka.VoteGetOwnedWordSetList: {},
		},
	})

	res = voters
	return
}
