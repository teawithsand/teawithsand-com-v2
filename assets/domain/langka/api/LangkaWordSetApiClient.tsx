import { RemoteWordSetCreateRequest, RemoteWordSetCreateResponse, RemoteWordSetDeleteRequest, RemoteWordSetDeleteResponse, RemoteWordSetEditRequest, RemoteWordSetEditResponse, RemoteWordSetGetOwnedListRequest, RemoteWordSetGetOwnedListResponse, RemoteWordSetGetPublicListRequest, RemoteWordSetGetPublicListResponse, RemoteWordSetGetPublicRequest, RemoteWordSetGetPublicResponse, RemoteWordSetGetSecretRequest, RemoteWordSetGetSecretResponse, RemoteWordSetPublishRequest, RemoteWordSetPublishResponse } from "@app/exported/typings/api";
import { SimpleHttpClient } from "@app/util/apiClient/client";

export default class LangkaWordSetApiClient {
    constructor(private readonly client: SimpleHttpClient) {
    }

    createWordSet = this.client.apiFetcher<RemoteWordSetCreateRequest, RemoteWordSetCreateResponse>({
        url: "/api/v1/langka/word-set",
        method: "POST",
    })

    editWordSet = this.client.apiFetcher<RemoteWordSetEditRequest, RemoteWordSetEditResponse>({
        url: "/api/v1/langka/word-set",
        method: "PATCH",
    })

    deleteWordSet = this.client.apiFetcher<RemoteWordSetDeleteRequest, RemoteWordSetDeleteResponse>({
        url: "/api/v1/langka/word-set",
        method: "DELETE",
    })

    publishWordSet = this.client.apiFetcher<RemoteWordSetPublishRequest, RemoteWordSetPublishResponse>({
        url: "/api/v1/langka/word-set/publish",
        method: "POST",
    })

    getPublicWordSet = this.client.apiSearchFetcher<RemoteWordSetGetPublicRequest, RemoteWordSetGetPublicResponse>({
        url: "/api/v1/langka/word-set/public",
        method: "GET",
    })

    getSecretWordSet = this.client.apiSearchFetcher<RemoteWordSetGetSecretRequest, RemoteWordSetGetSecretResponse>({
        url: "/api/v1/langka/word-set/public",
        method: "GET",
    })

    getPublicWordSetList = this.client.apiSearchFetcher<RemoteWordSetGetPublicListRequest, RemoteWordSetGetPublicListResponse>({
        url: "/api/v1/langka/word-set/list/public",
        method: "GET",
    })

    getOwnedWordSetList = this.client.apiSearchFetcher<RemoteWordSetGetOwnedListRequest, RemoteWordSetGetOwnedListResponse>({
        url: "/api/v1/langka/word-set/list/owned",
        method: "GET",
    })
}