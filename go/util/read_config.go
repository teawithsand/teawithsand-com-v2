package util

import "github.com/kelseyhightower/envconfig"

func ReadConfig(res interface{}) (err error) {
	err = envconfig.Process("TWSAPI", res)
	return
}
