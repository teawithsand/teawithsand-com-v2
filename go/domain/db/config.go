package db

type Config struct {
	MongoURL    string `required:"true" split_words:"true"`
	MongoDBName string `required:"true" split_words:"true"`
}
