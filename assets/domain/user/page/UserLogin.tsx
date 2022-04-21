import { commonHomeEndpoint, useEndpointNavigate } from "@app/domain/common/endpoints"
import { userSetDataAction } from "@app/domain/common/redux/user"
import { useApiClient } from "@app/util/apiClient/client"
import React from "react"
import { FormattedMessage } from "react-intl"
import { useDispatch } from "react-redux"
import UserApiClient from "../api/UserApiClient"
import LoginForm from "../form/LoginForm"

export default () => {
    const client = useApiClient()
    const userClient = new UserApiClient(client)
    const dispatch = useDispatch()
    const navigate = useEndpointNavigate()

    return <div className="page-user-login">
        <header className="page-user-login__header">
            <h1>
                <FormattedMessage id="user.page.login.title" />
            </h1>
        </header>
        <main className="page-user-login__form">
            <LoginForm
                onSubmit={async (formData) => {
                    const res = await userClient.login({
                        login: formData.username,
                        password: formData.password,
                    })

                    const { data } = res

                    dispatch(userSetDataAction({
                        id: data.user.id,
                        publicName: data.user.publicName,
                        token: data.token,
                    }))
                }}
                onSuccess={() => {
                    navigate(commonHomeEndpoint, null)
                }}
            />
        </main>
    </div>
}