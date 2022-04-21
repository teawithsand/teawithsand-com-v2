package dbutil

import (
	"errors"
	"reflect"
)

var ErrInvalidData = errors.New("util/dbutil: invalid data provided")

type TypeRegistryChecker struct {
	fields map[reflect.Type]struct{}
}

func (trc *TypeRegistryChecker) AddTypeWithData(data interface{}) *TypeRegistryChecker {
	return trc.AddType(reflect.TypeOf(data))
}

func (trc *TypeRegistryChecker) AddType(ty reflect.Type) *TypeRegistryChecker {
	trc.fields[ty] = struct{}{}
	return trc
}

func (trc *TypeRegistryChecker) Check(data interface{}) (err error) {
	if data == nil {
		err = ErrInvalidData
		return
	}
	if trc.fields == nil {
		err = ErrInvalidData
		return
	}

	ty := reflect.TypeOf(data)
	_, ok := trc.fields[ty]
	if !ok {
		err = ErrInvalidData
		return
	}

	return
}
