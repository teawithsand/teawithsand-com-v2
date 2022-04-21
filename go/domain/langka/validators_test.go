package langka_test

import (
	"testing"

	"github.com/teawithsand/webpage/domain/langka"
	"github.com/teawithsand/webpage/util/testutil"
)

func TestValidationsOK(t *testing.T) {
	ctx := testutil.Context(t)
	validators := langka.DefaultValidators

	err := validators.WordSetCreateDataValidator.Validate(ctx, langka.WordSetCreateEditData{
		Name:         "Valid word set name!",
		FromLanguage: "en-US",
		ToLanguage:   "pl-PL",
		WordTuples: []langka.WordTupleCreateData{
			{
				SourceWord:       "apple",
				DestinationWords: []string{"jabłko"},
			},
			{
				SourceWord:       "let",
				DestinationWords: []string{"pozwolić", "niech"},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
}
