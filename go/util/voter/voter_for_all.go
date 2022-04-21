package voter

import "context"

type ForAll struct{}

func (*ForAll) Supports(ctx context.Context, voting Voting) (ok bool, err error) {
	ok = true
	return
}
func (*ForAll) Vote(ctx context.Context, voting Voting) (err error) {
	return
}
