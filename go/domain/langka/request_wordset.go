package langka

import (
	"context"

	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util/cext"
	"github.com/teawithsand/webpage/util/dbutil"
	"github.com/teawithsand/webpage/util/slug"
	"go.mongodb.org/mongo-driver/bson"
)

type WordTupleCreateData struct {
	SourceWord       string   `json:"sourceWord"`
	DestinationWords []string `json:"destinationWords"`
	Description      string   `json:"description"`
}

func (data WordTupleCreateData) Create() (tpl WordTuple) {
	tpl.Description = data.Description
	tpl.DestinationWords = data.DestinationWords
	tpl.SourceWord = data.SourceWord
	return
}

type WordSetCreateEditData struct {
	Name         string `json:"name"`
	FromLanguage string `json:"fromLanguage"`
	ToLanguage   string `json:"toLanguage"`
	Description  string `json:"description"`

	WordTuples []WordTupleCreateData `json:"wordTuples"` // nil != empty slice here
}

func (data *WordSetCreateEditData) Create(ctx context.Context, sluggifier slug.Sluggifier) (res *WordSet, err error) {
	ws := WordSet{}
	ws.ID = db.NewID()
	ws.FromLanguage = data.FromLanguage
	ws.ToLanguage = data.ToLanguage
	ws.Name = data.Name
	ws.Description = data.Description
	ws.NameSlug, err = sluggifier.Slugify(ctx, data.Name)
	if err != nil {
		return
	}

	ws.Lifecycle.CreatedAt = cext.Now(ctx)
	ws.Lifecycle.LastUpdatedAt = cext.Now(ctx)

	for _, td := range data.WordTuples {
		ws.WordTuples = append(ws.WordTuples, td.Create())
	}

	res = &ws

	return
}

func (data *WordSetCreateEditData) MakeDBUpdate(ctx context.Context, sluggifier slug.Sluggifier, force bool) (patch interface{}, err error) {
	res := bson.M{}

	res["lifecycle"] = bson.M{
		"lastupdatedat": cext.Now(ctx),
	}

	if force || len(data.FromLanguage) > 0 {
		res["fromlanguage"] = data.FromLanguage
	}
	if force || len(data.ToLanguage) > 0 {
		res["toLanguage"] = data.ToLanguage
	}

	if force || len(data.Name) > 0 {
		res["name"] = data.Name
		res["nameslug"], err = sluggifier.Slugify(ctx, data.Name)
		if err != nil {
			return
		}
	}

	if force || data.WordTuples != nil {
		var tuples []WordTuple
		for _, td := range data.WordTuples {
			tuples = append(tuples, td.Create())
		}

		res["tuples"] = tuples
	}

	patch = bson.M{
		"$set": res,
	}

	return
}

type WordSetCreateRequest struct {
	WordSetCreateEditData
}

type WordSetCreateResponse struct {
	ID db.ID `json:"id"`
}

type WordSetEditRequest struct {
	ID db.ID `json:"id"`
	WordSetCreateEditData
}
type WordSetEditResponse struct {
	// Empty
}

type WordSetDeleteRequest struct {
	ID db.ID `json:"id"`
}
type WordSetDeleteResponse struct {
	// Empty
}

type WordSetPublishRequest struct {
	ID      db.ID `json:"id"`
	Publish bool  `json:"publish"`
}
type WordSetPublishResponse struct {
	// Empty
}

type WordSetGetSecretRequest struct {
	ID db.ID `json:"id"`
}
type WordSetGetSecretResponse struct {
	WordSetSecretProjection
}

type WordSetGetPublicRequest struct {
	ID db.ID `json:"id"`
}
type WordSetGetPublicResponse struct {
	WordSetPublicProjection
}

type WordSetSearchManyParams struct {
	OwnerID             db.ID  `json:"owner,omitempty" schema:"owner"`
	OwnerName           string `json:"ownerName,omitempty" schema:"ownerName"`
	WordSetName         string `json:"wordSetName,omitempty" schema:"wordSetName"`
	SourceLanguage      string `json:"sourceLanguage,omitempty" schema:"sourceLanguage"`
	DestinationLanguage string `json:"destinationLanguage,omitempty" schema:"destinationLanguage"`

	Order dbutil.OrderFields `json:"order,omitempty" schema:"order"`
}

func (sp *WordSetSearchManyParams) MakeDBQuery() interface{} {
	var results []interface{}

	if !sp.OwnerID.IsZero() {
		results = append(results, dbutil.NewIDQuery(sp.OwnerID))
	}

	if sp.OwnerName != "" {
		results = append(results, dbutil.NewEQQuery("owner.publicname", sp.OwnerName))
	}

	if sp.WordSetName != "" {
		results = append(results, dbutil.NewEQQuery("name", sp.WordSetName))
	}

	if sp.SourceLanguage != "" {
		results = append(results, dbutil.NewEQQuery("sourceLanguage", sp.SourceLanguage))
	}

	if sp.DestinationLanguage != "" {
		results = append(results, dbutil.NewEQQuery("destinationLanguage", sp.DestinationLanguage))
	}

	return dbutil.QueryAnd(results...)
}

type WordSetGetPublicListRequest struct {
	WordSetSearchManyParams
	dbutil.Pagination
}

type WordSetGetPublicListResponse struct {
	Entries []WordSetPublicSummaryProjection `json:"entries"`
}

type WordSetGetOwnedListRequest struct {
	WordSetSearchManyParams
	dbutil.Pagination
}

type WordSetGetOwnedListResponse struct {
	Entries []WordSetSecretSummaryProjection `json:"entries"`
}
