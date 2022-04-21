package langka

import (
	"context"

	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/dbutil"
	"github.com/teawithsand/webpage/util/slug"
	"github.com/teawithsand/webpage/util/validator"
	"github.com/teawithsand/webpage/util/voter"
)

type WordSetService struct {
	WordSetRepository WordSetRepository
	Checker           voter.Checker
	Sluggifier        slug.Sluggifier

	CreateWordSetDataValidator validator.Validator
}

func (svc *WordSetService) CreateWordSet(ctx context.Context, req WordSetCreateRequest) (res WordSetCreateResponse, err error) {
	user := user.GetUser(ctx)

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteCreateWordSet,
		Object:    nil,
	})
	if err != nil {
		return
	}

	if user == nil {
		panic("domain/langka: user may not be nil here")
	}

	err = svc.CreateWordSetDataValidator.Validate(ctx, req.WordSetCreateEditData)
	if err != nil {
		return
	}

	ws, err := req.Create(ctx, svc.Sluggifier)
	if err != nil {
		return
	}

	ws.Owner.Load(user)

	err = svc.WordSetRepository.Create(ctx, *ws)
	if db.IsUniqueViolatedError(err) {
		err = ErrWordSetNameInUse
		return
	} else if err != nil {
		return
	}

	res.ID = ws.ID

	return
}

func (svc *WordSetService) EditWordSet(ctx context.Context, req WordSetEditRequest) (res WordSetCreateResponse, err error) {
	var ws WordSet
	ok, err := svc.WordSetRepository.FindOne(ctx, dbutil.NewIDQuery(req.ID), &ws)
	if err != nil {
		return
	}

	if !ok {
		err = ErrWordSetNotFound
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteEditWordSet,
		Object:    &ws,
	})
	if err != nil {
		return
	}

	update, err := req.MakeDBUpdate(ctx, svc.Sluggifier, false)
	if err != nil {
		return
	}

	err = svc.WordSetRepository.UpdateOne(ctx, dbutil.NewIDQuery(req.ID), update)
	if db.IsUniqueViolatedError(err) {
		err = ErrWordSetNameInUse
		return
	} else if err != nil {
		return
	}

	return
}

func (svc *WordSetService) DeleteWordSet(ctx context.Context, req WordSetDeleteRequest) (res WordSetDeleteResponse, err error) {
	var ws WordSet
	ok, err := svc.WordSetRepository.FindOne(ctx, dbutil.NewIDQuery(req.ID), &ws)
	if err != nil {
		return
	}

	if !ok {
		err = ErrWordSetNotFound
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteDeleteWordSet,
		Object:    &ws,
	})
	if err != nil {
		return
	}

	err = svc.WordSetRepository.DeleteOne(ctx, dbutil.NewIDQuery(req.ID))
	if err != nil {
		return
	}

	return
}

func (svc *WordSetService) PublishWordSet(ctx context.Context, req WordSetPublishRequest) (res WordSetPublishRequest, err error) {
	var ws WordSet
	ok, err := svc.WordSetRepository.FindOne(ctx, dbutil.NewIDQuery(req.ID), &ws)
	if err != nil {
		return
	}

	if !ok {
		err = ErrWordSetNotFound
		return
	}

	if req.Publish {
		err = svc.Checker.Check(ctx, voter.Voting{
			Operation: VotePublishWordSet,
			Object:    &ws,
		})
		if err != nil {
			return
		}
	} else {
		err = svc.Checker.Check(ctx, voter.Voting{
			Operation: VoteUnpublishWordSet,
			Object:    &ws,
		})
		if err != nil {
			return
		}
	}

	var update interface{}
	if req.Publish {
		update = WordSetPublishUpdate(ctx)
	} else {
		update = WordSetUnpublishUpdate(ctx)
	}

	err = svc.WordSetRepository.UpdateOne(ctx, dbutil.NewIDQuery(req.ID), update)
	if err != nil {
		return
	}

	return
}

func (svc *WordSetService) GetSecretWordSet(ctx context.Context, req WordSetGetSecretRequest) (res WordSetGetSecretResponse, err error) {
	var ws WordSet
	ok, err := svc.WordSetRepository.FindOne(ctx, dbutil.NewIDQuery(req.ID), &ws)
	if err != nil {
		return
	}

	if !ok {
		err = ErrWordSetNotFound
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteGetSecretWordSet,
		Object:    &ws,
	})
	if err != nil {
		return
	}

	res.Load(&ws)
	return
}

func (svc *WordSetService) GetPublicWordSet(ctx context.Context, req WordSetGetPublicRequest) (res WordSetGetPublicResponse, err error) {
	var ws WordSet
	ok, err := svc.WordSetRepository.FindOne(ctx, dbutil.NewIDQuery(req.ID), &ws)
	if err != nil {
		return
	}

	if !ok {
		err = ErrWordSetNotFound
		return
	}

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteGetPublicWordSet,
		Object:    &ws,
	})
	if err != nil {
		return
	}

	res.Load(&ws)
	return
}

func (svc *WordSetService) GetPublicWordSetList(ctx context.Context, req WordSetGetPublicListRequest) (res WordSetGetPublicListResponse, err error) {
	cursor, err := svc.WordSetRepository.FindMany(ctx, dbutil.QueryAnd(
		req.MakeDBQuery(),
		WordSetPublishedQuery(true),
	), req.Order, req.Pagination.WithMaxLimit(30))
	if err != nil {
		return
	}
	defer cursor.Close()

	for cursor.Next(nil) {
		var ws WordSet
		err = cursor.Decode(&ws)
		if err != nil {
			return
		}

		var proj WordSetPublicSummaryProjection
		proj.Load(&ws)
		res.Entries = append(res.Entries, proj)
	}

	return
}

func (svc *WordSetService) GetOwnerdWordSetList(ctx context.Context, req WordSetGetOwnedListRequest) (res WordSetGetOwnedListResponse, err error) {
	u := user.GetUser(ctx)

	err = svc.Checker.Check(ctx, voter.Voting{
		Operation: VoteGetOwnedWordSetList,
		Object:    nil,
	})
	if err != nil {
		return
	}

	if u == nil {
		panic("user must not be nil here")
	}

	cursor, err := svc.WordSetRepository.FindMany(ctx, dbutil.QueryAnd(
		req.MakeDBQuery(),
		dbutil.NewEQQuery("owner._id", u.ID),
	), req.Order, req.Pagination.WithMaxLimit(30))
	if err != nil {
		return
	}
	defer cursor.Close()

	for cursor.Next(nil) {
		var ws WordSet
		err = cursor.Decode(&ws)
		if err != nil {
			return
		}

		var proj WordSetSecretSummaryProjection
		proj.Load(&ws)
		res.Entries = append(res.Entries, proj)
	}

	return
}
