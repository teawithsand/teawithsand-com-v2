import { useEndpointNavigate, userLoginEndpoint } from "@app/domain/common/endpoints"
import { useApiClient } from "@app/util/apiClient/client"
import React from "react"
import { FormattedMessage } from "react-intl"
import { useSearchParams } from "react-router-dom"
import UserApiClient from "../api/UserApiClient"
import ConfirmRegistrationForm from "../form/ConfirmRegistrationForm"

export default () => {
    const client = useApiClient()
    const userClient = new UserApiClient(client)
    const navigate = useEndpointNavigate()

    const [params, _] = useSearchParams({
        token: "",
    })
    const token = params.get("token") ?? ""

    if (token) {
        return <div className="page-user-confirm-registration">
            <header className="page-user-confirm-registration__header">
                <h1>
                    <FormattedMessage id="user.page.confirm_registration.title" />
                </h1>
            </header>
            <main className="page-user-confirm-registration__form">
                <ConfirmRegistrationForm
                    onSubmit={async (formData) => {
                        await userClient.confirmRegister({
                            password: formData.password,
                            token,
                        })
                    }}
                    onSuccess={() => {
                        navigate(userLoginEndpoint, null)
                    }}
                />
            </main>
        </div>
    } else {
        return <div className="page-user-confirm-registration">
            <header className="page-user-confirm-registration__header">
                <h1>
                    <FormattedMessage id="user.page.confirm_registration.title" />
                </h1>
            </header>
            <main className="page-user-confirm-registration__error">
                <h3>
                    <FormattedMessage id="user.page.confirm_registration.title" />
                </h3>
            </main>
        </div>
    }
}