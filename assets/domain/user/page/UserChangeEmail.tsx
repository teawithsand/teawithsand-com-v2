import ExplainedLoadingError from "@app/domain/common/component/ExplainedLoadingError"
import { useEndpointNavigate, userSecretProfileEndpoint } from "@app/domain/common/endpoints"
import { userDataSelector } from "@app/domain/common/redux/user"
import { useApiClient } from "@app/util/apiClient/client"
import React from "react"
import { FormattedMessage } from "react-intl"
import { useSelector } from "react-redux"
import { userMustLogInError } from "../api/error"
import UserApiClient from "../api/UserApiClient"
import ChangeEmailForm from "../form/ChangeEmailForm"
import ChangePasswordForm from "../form/ChangePasswordForm"

export default () => {
    const client = useApiClient()
    const userClient = new UserApiClient(client)
    const navigate = useEndpointNavigate()

    const userData = useSelector(userDataSelector)

    if (!userData) {
        return <ExplainedLoadingError
            error={userMustLogInError}
        />
    }

    return <div className="page-user-change-email">
        <header className="page-user-change-email__header">
            <h1>
                <FormattedMessage id="user.page.change_email.title" />
            </h1>
        </header>
        <main className="page-user-change-email__form">
            <ChangeEmailForm
                onSubmit={async (formData) => {
                    await userClient.changeEmail({
                        email: formData.email,
                        captchaResponse: formData.captchaResponse,
                    })
                }}
                onSuccess={() => {
                    navigate(userSecretProfileEndpoint, {
                        id: userData.id,
                    })
                }}
            />
        </main>
    </div>
}