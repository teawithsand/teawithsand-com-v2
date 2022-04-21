import { useEndpointNavigate, userLoginEndpoint } from "@app/domain/common/endpoints"
import { useApiClient } from "@app/util/apiClient/client"
import React from "react"
import { FormattedMessage } from "react-intl"
import UserApiClient from "../api/UserApiClient"
import RegisterForm from "../form/RegisterForm"

export default () => {
    const client = useApiClient()
    const userClient = new UserApiClient(client)

    const epNavigate = useEndpointNavigate()


    return <div className="page-user-register">
        <header className="page-user-register__header">
            <h1>
                <FormattedMessage id="user.page.register.title" />
            </h1>
        </header>
        <main className="page-user-register__form">
            <RegisterForm
                onSubmit={async (data) => {
                    await userClient.register(data)
                }}
                onSuccess={() => {
                    epNavigate(userLoginEndpoint, null)
                }}
            />
        </main>
    </div>
}