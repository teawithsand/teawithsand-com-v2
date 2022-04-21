package remexpl

import (
	"sync"

	"github.com/teawithsand/webpage/util/explain"
)

func getExplainers() explain.Explainers {
	var explainers explain.Explainers

	explainers, err := registerUserExplainers(explainers)
	if err != nil {
		panic(err)
	}

	explainers, err = registerLangkaExplainers(explainers)
	if err != nil {
		panic(err)
	}

	return explainers
}

var once sync.Once

func InitExplainers() {
	once.Do(func() {
		for _, e := range getExplainers() {
			explain.AddGlobal(e)
		}
	})
}
