import { RemoteRegisterRequest, RemoteLoginRequest, RemoteLoginResponse, RemoteInvalidateTokensRequest, RemoteInvalidateTokensResponse, RemoteConfirmRegistrationRequest, RemoteRegisterResponse, RemoteConfirmRegistrationResponse, RemoteResetPasswordRequest, RemoteResetPasswordResponse, RemoteInitResetPasswordRequest, RemoteInitResetPasswordResponse, RemoteChangeEmailRequest, RemoteChangeEmailResponse, RemoteConfirmEmailRequest, RemoteConfirmEmailResponse, RemoteGetUserSecretProfileRequest, RemoteGetUserSecretProfileResponse, RemoteChangePasswordRequest, RemoteChangePasswordResponse } from "@app/exported/typings/api";
import { ApiResponse, SimpleHttpClient } from "@app/util/apiClient/client";

export default class UserApiClient {
    constructor(private readonly client: SimpleHttpClient) {
    }

    private makeRequestFac = <RQ, RS>(path: string): (req: RQ) => Promise<ApiResponse<RS>> => {
        return this.client.apiFetcher({
            url: path,
            method: "POST",
        })
    }

    register = this.makeRequestFac<RemoteRegisterRequest, RemoteRegisterResponse>("/api/v1/user/register/init")
    confirmRegister = this.makeRequestFac<RemoteConfirmRegistrationRequest, RemoteConfirmRegistrationResponse>("/api/v1/user/register/confirm")

    login = this.makeRequestFac<RemoteLoginRequest, RemoteLoginResponse>("/api/v1/user/login")
    
    invalidateTokens = this.makeRequestFac<RemoteInvalidateTokensRequest, RemoteInvalidateTokensResponse>("/api/v1/user/invalidate-tokens")

    initResetPassword = this.makeRequestFac<RemoteInitResetPasswordRequest, RemoteInitResetPasswordResponse>("/api/v1/user/reset-password/init")
    resetPassword = this.makeRequestFac<RemoteResetPasswordRequest, RemoteResetPasswordResponse>("/api/v1/user/reset-password/finalize")

    changePassword = this.client.apiFetcher<
        RemoteChangePasswordRequest,
        RemoteChangePasswordResponse
    >({ url: "/api/v1/user/change-password" })

    changeEmail = this.makeRequestFac<RemoteChangeEmailRequest, RemoteChangeEmailResponse>("/api/v1/user/change-email/init")
    confirmEmail = this.makeRequestFac<RemoteConfirmEmailRequest, RemoteConfirmEmailResponse>("/api/v1/user/change-email/finalize")

    getSecretProfile = this.client.apiSearchFetcher<
        RemoteGetUserSecretProfileRequest,
        RemoteGetUserSecretProfileResponse
    >({ url: "/api/v1/user/profile/secret" })
}