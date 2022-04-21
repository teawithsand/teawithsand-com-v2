package voter

import "context"

type Operation string

type Voting struct {
	Operation Operation
	Object    interface{}
}

type Voter interface {
	// Nil error == that voter grants access.
	Supports(ctx context.Context, voting Voting) (ok bool, err error)
	Vote(ctx context.Context, voting Voting) (err error)
}

// Checker checks if permission is granted or not. It consists of multiple voters.
type Checker interface {
	Check(ctx context.Context, voting Voting) (err error)
}
