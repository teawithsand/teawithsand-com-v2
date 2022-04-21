package dbutil

import (
	"context"
	"errors"
	"io"

	"github.com/teawithsand/webpage/domain/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	FindOne(ctx context.Context, query interface{}, res interface{}) (ok bool, err error)
	FindMany(
		ctx context.Context,
		query interface{},
		orderFields OrderFields,
		pagination Pagination,
	) (cursor Cursor, err error)

	UpdateOne(ctx context.Context, query interface{}, update interface{}) (err error)

	Create(ctx context.Context, data interface{}) (err error)

	DeleteOne(ctx context.Context, query interface{}) (err error)
	DeleteMany(ctx context.Context, query interface{}) (err error)
}

type Cursor interface {
	io.Closer

	Next(ctx context.Context) bool
	Decode(res interface{}) (err error)
}

type RepositoryImpl struct {
	Collection db.Collection

	DestinationChecker func(res interface{}) (err error)
	CreateDataChecker  func(data interface{}) (err error)
	QueryChecker       func(query interface{}) (err error)
	OrderFieldsChecker func(fields OrderFields) (err error)
}

type cursorImpl struct {
	repo   *RepositoryImpl
	ctx    context.Context
	cursor *mongo.Cursor
}

func (c *cursorImpl) Next(ctx context.Context) bool {
	if ctx == nil {
		ctx = c.ctx
	}
	return c.cursor.Next(ctx)
}

func (c *cursorImpl) Decode(res interface{}) (err error) {
	if c.repo.DestinationChecker != nil {
		err = c.repo.DestinationChecker(res)
		if err != nil {
			return
		}
	}

	err = c.cursor.Decode(res)
	if err != nil {
		return
	}
	return
}

func (c *cursorImpl) Close() (err error) {
	return c.cursor.Close(c.ctx)
}

func (r *RepositoryImpl) FindMany(
	ctx context.Context,
	query interface{},
	orderFields OrderFields,
	pagination Pagination,
) (cursor Cursor, err error) {
	if r.OrderFieldsChecker != nil {
		err = r.OrderFieldsChecker(orderFields)
		if err != nil {
			return
		}
	}

	idQuery, ok := query.(db.ID)
	if ok {
		query = NewIDQuery(idQuery)
	}

	if r.QueryChecker != nil {
		err = r.QueryChecker(query)
		if err != nil {
			return
		}
	}

	options := &options.FindOptions{}
	options = options.SetSort(orderFields.GetFields())
	options = options.SetLimit(int64(pagination.Limit))
	options = options.SetSkip(int64(pagination.Offset))

	rawCursor, err := r.Collection.Find(ctx, query, options)
	if err != nil {
		return
	}

	cursor = &cursorImpl{
		ctx:    ctx,
		cursor: rawCursor,
		repo:   r,
	}
	return
}

func (r *RepositoryImpl) FindOne(ctx context.Context, query interface{}, res interface{}) (ok bool, err error) {
	if res != nil && r.DestinationChecker != nil {
		err = r.DestinationChecker(res)
		if err != nil {
			return
		}
	}

	idQuery, ok := query.(db.ID)
	if ok {
		query = NewIDQuery(idQuery)
	}

	if r.QueryChecker != nil {
		err = r.QueryChecker(query)
		if err != nil {
			return
		}
	}

	tmp := r.Collection.FindOne(ctx, query)

	if res != nil {
		err = tmp.Decode(res)
	} else {
		err = tmp.Err()
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		ok = false
		err = nil
		return
	} else if err != nil {
		return
	}

	ok = true
	return
}

func (r *RepositoryImpl) UpdateOne(ctx context.Context, query interface{}, update interface{}) (err error) {
	_, err = r.Collection.UpdateOne(ctx, query, update)
	return
}
func (r *RepositoryImpl) Create(ctx context.Context, data interface{}) (err error) {
	if r.CreateDataChecker != nil {
		err = r.CreateDataChecker(data)
		if err != nil {
			return
		}
	}

	_, err = r.Collection.InsertOne(ctx, data)
	return
}

func (r *RepositoryImpl) DeleteOne(ctx context.Context, query interface{}) (err error) {
	if r.QueryChecker != nil {
		err = r.QueryChecker(query)
		if err != nil {
			return
		}
	}

	_, err = r.Collection.DeleteOne(ctx, query)
	return
}

func (r *RepositoryImpl) DeleteMany(ctx context.Context, query interface{}) (err error) {
	if r.QueryChecker != nil {
		err = r.QueryChecker(query)
		if err != nil {
			return
		}
	}

	_, err = r.Collection.DeleteOne(ctx, query)
	return
}
