package voter

import "context"

// VoterFunc is function, which supports voting on any operation.
type VoterFunc func(ctx context.Context, voting Voting) (err error)

func (f VoterFunc) Supports(ctx context.Context, voting Voting) (ok bool, err error) {
	ok = true
	return
}
func (f VoterFunc) Vote(ctx context.Context, voting Voting) (err error) {
	return f(ctx, voting)
}
