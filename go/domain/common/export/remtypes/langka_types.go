package remtypes

import "github.com/teawithsand/webpage/domain/langka"

func registerLangkaTypes() {
	addType(langka.WordTupleCreateData{})
	addType(langka.WordSetCreateEditData{})

	addType(langka.WordSetPublicProjection{})
	addType(langka.WordSetSecretProjection{})
	addType(langka.WordTuple{})

	addType(langka.WordSetCreateRequest{})
	addType(langka.WordSetCreateResponse{})

	addType(langka.WordSetEditRequest{})
	addType(langka.WordSetEditResponse{})

	addType(langka.WordSetDeleteRequest{})
	addType(langka.WordSetDeleteResponse{})

	addType(langka.WordSetPublishRequest{})
	addType(langka.WordSetPublishResponse{})

	addType(langka.WordSetGetPublicRequest{})
	addType(langka.WordSetGetPublicResponse{})

	addType(langka.WordSetGetSecretRequest{})
	addType(langka.WordSetGetSecretResponse{})

	addType(langka.WordSetGetPublicListRequest{})
	addType(langka.WordSetGetPublicListResponse{})

	addType(langka.WordSetGetOwnedListRequest{})
	addType(langka.WordSetGetOwnedListResponse{})
}
