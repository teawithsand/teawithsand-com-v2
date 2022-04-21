package remtypes

import (
	"time"

	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util/dbutil"
	"github.com/teawithsand/webpage/util/typescript"
)

var converter = typescript.New().WithPrefix("Remote").WithInterface(true)

func GetConverter() *typescript.Converter {
	return converter
}

func addType(ty interface{}) {
	converter = converter.Add(ty)
}

func addVar(name string, value interface{}) {
	converter = converter.AddVariable(name, value)
}

func init() {
	converter = converter.
		ManageType(time.Time{}, typescript.TypeOptions{TSType: "Date", TSTransform: "new Date(__VALUE__)"}).
		ManageType(db.ID{}, typescript.TypeOptions{
			TSType:      "string",
			TSTransform: "__VALUE__",
		}).
		ManageType(dbutil.OrderField{}, typescript.TypeOptions{TSType: "string"}).
		ManageType(dbutil.OrderFields{}, typescript.TypeOptions{TSType: "string"})

	registerUserTypes()
	registerLangkaTypes()
}
