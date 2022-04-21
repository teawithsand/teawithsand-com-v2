package env

import "github.com/teawithsand/webpage/util/testutil"

type Config struct {
	ENV string
}

func ReadConfig() Config {
	cfg := Config{
		ENV: "dev",
	}

	if testutil.IsTest() {
		cfg.ENV = "test"
	}

	return cfg
}
