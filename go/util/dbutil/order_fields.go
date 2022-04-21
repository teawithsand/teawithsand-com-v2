package dbutil

import (
	"bytes"
	"encoding"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

var ErrInvalidOrderFields = errors.New("domain/db: invalid order fields provided")
var ErrFieldTwice = errors.New("domain/db: some field was found twice")

type OrderSchema struct {
	AliasToField map[string]string
}

// TODO(teawithsand): validate schema with structure it's supposed to operate on

func (schema *OrderSchema) AddField(userAlias, dbName string) *OrderSchema {
	if schema.AliasToField == nil {
		schema.AliasToField = make(map[string]string)
	}

	if len(dbName) == 0 {
		dbName = userAlias
	}

	schema.AliasToField[userAlias] = dbName

	return schema
}

func (schema *OrderSchema) Validate(fields OrderFields) (err error) {
	if len(schema.AliasToField) == 0 {
		if len(fields) != 0 {
			err = ErrInvalidOrderFields
			return
		}
		return
	}
	usedMap := make(map[string]struct{})

	for _, f := range fields {
		_, ok := usedMap[f.Name]
		if ok {
			err = ErrFieldTwice
			return
		}

		_, ok = schema.AliasToField[f.Name]
		if !ok {
			err = ErrInvalidOrderFields
			return
		}

		usedMap[f.Name] = struct{}{}
	}

	return
}

// Processes fields, and applies aliases
func (schema *OrderSchema) Process(fields OrderFields) (res OrderFields) {
	res = make(OrderFields, 0, len(fields)/2)
	if schema.AliasToField == nil {
		return
	}
	usedMap := make(map[string]struct{})
	for _, aliasField := range fields {
		_, ok := usedMap[aliasField.Name]
		if ok {
			continue
		}

		dbName, ok := schema.AliasToField[aliasField.Name]
		if ok {
			res = append(res, OrderField{
				Name: dbName,
				Desc: aliasField.Desc,
			})
		}

		usedMap[aliasField.Name] = struct{}{}
	}
	return
}

type OrderField struct {
	Name string
	Desc bool
}

var _ encoding.TextMarshaler = &OrderField{}
var _ encoding.TextUnmarshaler = &OrderField{}

var ErrInvalidFirstChar = errors.New("util/dbutil: OrderField first char must be + or -")

func (f *OrderField) MarshalText() ([]byte, error) {
	res := make([]byte, len(f.Name)+1)
	if f.Desc {
		res[0] = '+'
	} else {
		res[0] = '-'
	}

	return res, nil
}

func (f *OrderField) UnmarshalText(data []byte) (err error) {
	if len(data) == 0 {
		err = ErrInvalidFirstChar
		return
	}

	if data[0] == '+' {
		f.Desc = false
	} else if data[0] == '-' {
		f.Desc = true
	} else {
		err = ErrInvalidFirstChar
		return
	}

	f.Name = string(data[1:])
	return
}

type OrderFields []OrderField

var _ encoding.TextMarshaler = &OrderFields{}
var _ encoding.TextUnmarshaler = &OrderFields{}

func (fields OrderFields) GetFields() bson.D {
	doc := bson.D{}
	for _, field := range fields {
		order := 1
		if field.Desc {
			order = -1
		}

		doc = append(doc, bson.E{
			Key:   field.Name,
			Value: order,
		})
	}
	return doc
}

func (fields *OrderFields) MarshalText() (res []byte, err error) {
	if fields == nil {
		return
	}
	for _, v := range *fields {
		var fieldRes []byte
		fieldRes, err = v.MarshalText()
		if err != nil {
			return
		}
		res = append(res, []byte(" ")...)
		res = append(res, fieldRes...)
	}

	return
}

func (fields *OrderFields) UnmarshalText(data []byte) (err error) {
	entries := bytes.Split(data, []byte(" "))
	for _, e := range entries {
		var f OrderField
		e = bytes.Trim(e, "\n\t\v ")
		err = f.UnmarshalText(e)
		if err != nil {
			return
		}
		*fields = append(*fields, f)
	}
	return
}
