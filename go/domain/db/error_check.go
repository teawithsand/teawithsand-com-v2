package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func IsUniqueViolatedError(err error) bool {
	if err == nil {
		return false
	}

	const code = 11000

	switch v := err.(type) {
	case mongo.WriteException:
		return v.HasErrorCode(code)
	case *mongo.WriteException:
		return v.HasErrorCode(code)
	case mongo.BulkWriteException:
		return v.HasErrorCode(code)
	case *mongo.BulkWriteException:
		return v.HasErrorCode(code)

	case mongo.WriteError:
		return v.Code == code
	case *mongo.WriteError:
		return v.Code == code
	case mongo.BulkWriteError:
		return v.Code == code
	case *mongo.BulkWriteError:
		return v.Code == code
	default:
		return false
	}
}
