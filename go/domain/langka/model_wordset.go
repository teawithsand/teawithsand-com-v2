package langka

import (
	"context"
	"time"

	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/cext"
	"github.com/teawithsand/webpage/util/dbutil"
	"go.mongodb.org/mongo-driver/bson"
)

type WordTuple struct {
	SourceWord       string   `json:"sourceWord"`
	DestinationWords []string `json:"destinationWords"`
	Description      string   `json:"description"`
}

type WordSetLifecycle struct {
	CreatedAt     time.Time
	PublishedAt   time.Time
	LastUpdatedAt time.Time
}

type WordSet struct {
	ID          db.ID `bson:"_id"`
	Owner       user.UserReference
	Name        string
	NameSlug    string
	Description string

	FromLanguage string
	ToLanguage   string

	WordTuples []WordTuple
	Lifecycle  WordSetLifecycle
}

type WordSetPublicProjection struct {
	ID    db.ID                        `bson:"_id"`
	Owner user.UserReferenceProjection `json:"owner"`

	Description  string `json:"description"`
	Name         string `json:"name"`
	FromLanguage string `json:"fromLanguage"`
	ToLanguage   string `json:"toLanguage"`

	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`

	WordTuples []WordTuple `json:"wordTuples"` // word tuple can be sent as-is
}

func (proj *WordSetPublicProjection) Load(ws *WordSet) {
	proj.Owner.Load(&ws.Owner)

	proj.ID = ws.ID
	proj.Description = ws.Description
	proj.Name = ws.Name
	proj.FromLanguage = ws.FromLanguage
	proj.ToLanguage = ws.ToLanguage

	proj.WordTuples = ws.WordTuples
}

type WordSetSecretProjection struct {
	WordSetPublicProjection
	PublishedAt *time.Time `json:"publishedAt"`
}

func (proj *WordSetSecretProjection) Load(ws *WordSet) {
	proj.WordSetPublicProjection.Load(ws)
	if !ws.Lifecycle.PublishedAt.IsZero() {
		proj.PublishedAt = &ws.Lifecycle.PublishedAt
	}
}

type WordSetPublicSummaryProjection struct {
	ID    db.ID                        `json:"id"`
	Owner user.UserReferenceProjection `json:"owner"`

	Name         string `json:"name"`
	FromLanguage string `json:"fromLanguage"`
	ToLanguage   string `json:"toLanguage"`

	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`

	WordTupleCount int `json:"wordTupleCount"`
}

func (proj *WordSetPublicSummaryProjection) Load(ws *WordSet) {
	proj.Owner.Load(&ws.Owner)

	proj.ID = ws.ID
	proj.Name = ws.Name
	proj.FromLanguage = ws.FromLanguage
	proj.ToLanguage = ws.ToLanguage

	proj.CreatedAt = ws.Lifecycle.CreatedAt
	proj.LastUpdatedAt = ws.Lifecycle.LastUpdatedAt

	proj.WordTupleCount = len(ws.WordTuples)
}

type WordSetSecretSummaryProjection struct {
	WordSetPublicSummaryProjection
	PublishedAt *time.Time `json:"publishedAt"`
}

func (proj *WordSetSecretSummaryProjection) Load(ws *WordSet) {
	proj.WordSetPublicSummaryProjection.Load(ws)
	if !ws.Lifecycle.PublishedAt.IsZero() {
		proj.PublishedAt = &ws.Lifecycle.PublishedAt
	}
}

func WordSetPublishUpdate(ctx context.Context) interface{} {
	return bson.M{
		"lifecycle": bson.M{
			"publishedat": cext.Now(ctx),
		},
	}
}

func WordSetUnpublishUpdate(ctx context.Context) interface{} {
	return bson.M{
		"lifecycle": bson.M{
			"publishedat": time.Time{},
		},
	}
}

func WordSetPublishedQuery(published bool) interface{} {
	query := dbutil.NewORQuery(
		dbutil.NewExistsQuery("lifecycle.publishedat", true),    // not null/exists
		dbutil.NewNEQuery("lifecycle.publishedat", time.Time{}), // is not equal to empty time
	) // then published
	if !published {
		query = dbutil.NewNotQuery(query)
	}
	return query
}
