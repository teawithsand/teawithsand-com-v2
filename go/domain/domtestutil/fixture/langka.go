package fixture

import (
	"time"

	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/langka"
)

var MockWordSetOne = langka.WordSet{
	ID:           db.NewID(),
	Owner:        MockUserOne.GetReference(),
	Name:         "Word set one",
	NameSlug:     "word-set-one-123456",
	FromLanguage: "en-US",
	ToLanguage:   "en-GB",
	WordTuples: []langka.WordTuple{
		{
			SourceWord:       "apple",
			DestinationWords: []string{"apple"},
			Description:      "What did you expect? It's the same language...",
		},
		{
			SourceWord:       "flashlight",
			DestinationWords: []string{"torch"},
			Description:      "With minor differences ofc",
		},
	},
	Lifecycle: langka.WordSetLifecycle{
		CreatedAt: time.Now(),
	},
}
