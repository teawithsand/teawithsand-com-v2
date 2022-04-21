package langka_test

import (
	"context"
	"testing"

	"github.com/teawithsand/webpage/domain/domtestutil"
	"github.com/teawithsand/webpage/domain/domtestutil/fixture"
	"github.com/teawithsand/webpage/domain/langka"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/dbutil"
	"github.com/teawithsand/webpage/util/testutil"
)

func TestWordSetService_CreateWordSet(t *testing.T) {
	testutil.SetTestENV()

	tdi := domtestutil.MustConstructTestDI()
	defer tdi.Close()

	ctx := context.Background()
	ctx, closer, err := tdi.DB.MakeSession(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	defer closer.Close()

	err = fixture.Apply(ctx, tdi.DB)
	if err != nil {
		t.Error(err)
		return
	}

	ctx = user.PutUser(ctx, &fixture.MockUserOne)

	res, err := tdi.LangkaWordSetService.CreateWordSet(ctx, langka.WordSetCreateRequest{
		WordSetCreateEditData: langka.WordSetCreateEditData{
			Name:         "asdf word set",
			FromLanguage: "en-US",
			ToLanguage:   "en-GB",

			WordTuples: []langka.WordTupleCreateData{
				{
					SourceWord:       "apple",
					DestinationWords: []string{"apple", "it's same language in the end"},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	var ws langka.WordSet
	ok, err := tdi.LangkaWordSetRepository.FindOne(ctx, dbutil.NewIDQuery(res.ID), &ws)
	if err != nil {
		t.Error(err)
		return
	}

	if !ok {
		t.Error("expected newly created ws to exist")
		return
	}

}

func TestWordSetService_EditWordSet(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		testutil.SetTestENV()

		tdi := domtestutil.MustConstructTestDI()
		defer tdi.Close()

		ctx := context.Background()
		ctx, closer, err := tdi.DB.MakeSession(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		defer closer.Close()

		err = fixture.Apply(ctx, tdi.DB)
		if err != nil {
			t.Error(err)
			return
		}

		const newName = "new name for ws one"

		ctx = user.PutUser(ctx, &fixture.MockUserOne)

		_, err = tdi.LangkaWordSetService.EditWordSet(ctx, langka.WordSetEditRequest{
			ID: fixture.MockWordSetOne.ID,
			WordSetCreateEditData: langka.WordSetCreateEditData{
				Name: newName,
			},
		})

		if err != nil {
			t.Error(err)
			return
		}

		var ws langka.WordSet
		ok, err := tdi.LangkaWordSetRepository.FindOne(ctx, dbutil.NewIDQuery(fixture.MockWordSetOne.ID), &ws)
		if err != nil {
			t.Error(err)
			return
		}

		if !ok {
			t.Error("expected edited ws to exist")
			return
		}

		if ws.Name != newName {
			t.Error("update filed")
			return
		}
	})

	t.Run("access_denied_invalid_user", func(t *testing.T) {
		testutil.SetTestENV()

		tdi := domtestutil.MustConstructTestDI()
		defer tdi.Close()

		ctx := context.Background()
		ctx, closer, err := tdi.DB.MakeSession(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		defer closer.Close()

		err = fixture.Apply(ctx, tdi.DB)
		if err != nil {
			t.Error(err)
			return
		}

		const newName = "new name for ws one"

		ctx = user.PutUser(ctx, &fixture.MockUserTwo)

		_, err = tdi.LangkaWordSetService.EditWordSet(ctx, langka.WordSetEditRequest{
			ID: fixture.MockWordSetOne.ID,
			WordSetCreateEditData: langka.WordSetCreateEditData{
				Name: newName,
			},
		})

		if err == nil {
			t.Error("expected error here")
			return
		}
	})
}
