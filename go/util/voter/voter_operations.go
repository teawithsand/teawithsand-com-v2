package voter

import "context"

// OperationsVoter is voter ,which checks support using set of operations given.
type OperationsVoter struct {
	Voter
	Operations map[Operation]struct{}
}

func (ov *OperationsVoter) Supports(ctx context.Context, voting Voting) (ok bool, err error) {
	if ov.Operations == nil {
		return
	}
	_, ok = ov.Operations[voting.Operation]
	if !ok {
		return
	}

	return ov.Voter.Supports(ctx, voting)
}
