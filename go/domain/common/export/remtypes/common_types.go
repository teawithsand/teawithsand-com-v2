package remtypes

import "github.com/teawithsand/webpage/util/dbutil"

func registerCommonTypes() {
	addType(dbutil.Pagination{})
}
