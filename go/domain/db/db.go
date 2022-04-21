package db

import (
	"context"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TypedDB struct {
	Database *mongo.Database
	Client   *mongo.Client

	Users         Collection
	Registrations Collection
	WordSets      Collection
}

type fnCloser func() error

func (f fnCloser) Close() error {
	return f()
}

func (tdb *TypedDB) MakeSession(ctx context.Context) (sessionContext context.Context, close io.Closer, err error) {
	session, err := tdb.Client.StartSession()
	if err != nil {
		return
	}
	sessionContext = mongo.NewSessionContext(ctx, session)

	close = fnCloser(func() error {
		session.EndSession(ctx)
		return nil
	})
	return
}

func (tdb *TypedDB) Clear(ctx context.Context) (err error) {
	_, err = tdb.Registrations.DeleteMany(ctx, bson.M{})
	if err != nil {
		return
	}

	_, err = tdb.Users.DeleteMany(ctx, bson.M{})
	if err != nil {
		return
	}

	_, err = tdb.WordSets.DeleteMany(ctx, bson.M{})
	if err != nil {
		return
	}

	return
}

type ID = primitive.ObjectID
type Collection = *mongo.Collection

func NewID() ID {
	return primitive.NewObjectID()
}

func MakeDB(ctx context.Context, cfg Config) (tdb *TypedDB, err error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURL))
	if err != nil {
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}

	rdb := client.Database(cfg.MongoDBName)

	tdb, err = MakeTypedDB(ctx, rdb)
	if err != nil {
		return
	}

	err = SetupDB(ctx, tdb)
	if err != nil {
		return
	}

	return
}

func MakeTypedDB(ctx context.Context, rdb *mongo.Database) (tdb *TypedDB, err error) {
	tdb = &TypedDB{
		Database:      rdb,
		Client:        rdb.Client(),
		Users:         rdb.Collection("users"),
		Registrations: rdb.Collection("registrations"),
		WordSets:      rdb.Collection("wordsets"),
	}
	return
}

// Note: this is done during connecting to db by default.
// No need to call it again.
func SetupDB(ctx context.Context, db *TypedDB) (err error) {
	userExpireAfter := time.Hour * 8

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	_, err = db.Registrations.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"createdat", 1},
			},
			Options: options.Index().
				SetExpireAfterSeconds(int32(userExpireAfter.Seconds())),
		},
		{
			Keys: bson.D{
				{"publicname", 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}, opts)
	if err != nil {
		return
	}

	_, err = db.Users.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"publicname", 1},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{"email.email", 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}, opts)

	if err != nil {
		return
	}

	_, err = db.WordSets.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"owner._id", 1},
			},
			Options: options.Index(),
		},
		{
			Keys: bson.D{
				{"name", 1},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{"nameslug", 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}, opts)

	if err != nil {
		return
	}

	return
}
