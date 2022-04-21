package webapp

type Config struct {
	ListenAddress string `required:"true" split_words:"true"`
}
