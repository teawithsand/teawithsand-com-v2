import ExplainedLoadingError from "@app/domain/common/component/ExplainedLoadingError"
import { useEndpointNavigate, userLoginEndpoint } from "@app/domain/common/endpoints"
import { unsetUnsetDataAction, userDataSelector } from "@app/domain/common/redux/user"
import { useApiClient } from "@app/util/apiClient/client"
import React from "react"
import { FormattedMessage } from "react-intl"
import { useDispatch, useSelector } from "react-redux"
import { userMustLogInError } from "../api/error"
import UserApiClient from "../api/UserApiClient"
import ChangePasswordForm from "../form/ChangePasswordForm"

export default () => {
    const client = useApiClient()
    const userClient = new UserApiClient(client)
    const navigate = useEndpointNavigate()

    const dispatch = useDispatch()
    const userData = useSelector(userDataSelector)

    if (!userData) {
        return <ExplainedLoadingError
            error={userMustLogInError}
        />
    }

    return <div className="page-user-change-password">
        <header className="page-user-change-password__header">
            <h1>
                <FormattedMessage id="user.page.change_password.title" />
            </h1>
        </header>
        <main className="page-user-change-password__form">
            <ChangePasswordForm
                onSubmit={async (formData) => {
                    await userClient.changePassword({
                        password: formData.password,
                    })
                }}
                onSuccess={() => {
                    dispatch(unsetUnsetDataAction())
                    navigate(userLoginEndpoint, null)
                }}
            />
        </main>
    </div>
}