package util

import "github.com/sarulabs/di"

type Builder = *di.Builder
type DI = di.Container
type Def = di.Def

func NewDIBuilder() Builder {
	builder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}
	return builder
}
