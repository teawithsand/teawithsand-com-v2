package user

import "github.com/teawithsand/webpage/domain/db"

type GetUserSecretProfileRequest struct {
	ID db.ID `json:"id"`
}

type GetUserSecretProfileResponse struct {
	SecretUserProjection
}
