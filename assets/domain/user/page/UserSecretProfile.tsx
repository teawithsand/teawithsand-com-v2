import LoadingSpinner from "@app/domain/common/component/LoadingSpinner"
import { useApiClient } from "@app/util/apiClient/client"
import { useRemoteData } from "@app/util/react/hook/useRemoteData"
import { useSearchParams } from "react-router-dom"
import UserApiClient from "../api/UserApiClient"

import React from "react"
import LoadingError from "@app/domain/common/component/LoadingError"
import { FormattedMessage, useIntl } from "react-intl"
import { defaultDateTimeFormatProps } from "@app/util/i18n/options"
import { getEndpointPath, userChangeEmailEndpoint, userChangePasswordEndpoint, userDeleteAccountEndpoint } from "@app/domain/common/endpoints"
import { Link } from "react-router-dom"
export default () => {
    const [params] = useSearchParams({
        id: "",
    })

    const id = params.get("id")

    const client = useApiClient()
    const userClient = new UserApiClient(client)
    const intl = useIntl()

    const rdl = useRemoteData({
        loader: async () => {
            return await userClient.getSecretProfile({ id })
        }
    }, [id])


    if (rdl.type === "pending") {
        return <LoadingSpinner />
    } else if (rdl.type === "error") {
        return <LoadingError error={rdl.error} />
    }

    const data = rdl.data.data

    return <div className="user-profile">
        <header className="user-profile__header">
            <h1>
                <FormattedMessage id="user.profile.secret.title" values={{ publicName: data.publicName }} />
            </h1>
        </header>
        <main className="user-profile__properties">
            <div>
                <FormattedMessage id="user.profile.secret.public_name" values={{
                    publicName: data.publicName
                }} />
            </div>
            <div>
                <FormattedMessage id="user.profile.secret.registered_at" values={{
                    registeredAt: intl.formatDate(new Date(data.registeredAt), {
                        ...defaultDateTimeFormatProps
                    }),
                }} />
            </div>
            <div>
                <FormattedMessage id="user.profile.secret.email" values={{ email: data.email }} />
            </div>
            {data.emailConfirmedAt ? <div>
                <FormattedMessage id="user.profile.secret.email_confirmed_at" values={{
                    confirmedAt: intl.formatDate(new Date(data.registeredAt), {
                        ...defaultDateTimeFormatProps
                    }),
                }} />
            </div> : <div>
                <h6><FormattedMessage id="user.profile.secret.email_not_confirmed" /></h6>
                <span>
                    <FormattedMessage id="user.profile.secret.email_not_confirmed.description" />
                </span>
            </div>}

        </main>
        <section className="user-profile__actions">
            <Link to={getEndpointPath(userDeleteAccountEndpoint, null)} className="user-profile__delete-btn">
                <FormattedMessage id="user.profile.delete_account" />
            </Link>

            <Link to={getEndpointPath(userChangeEmailEndpoint, null)} className="user-profile__change-email-btn">
                <FormattedMessage id="user.profile.change_email" />
            </Link>

            <Link to={getEndpointPath(userChangePasswordEndpoint, null)} className="user-profile__change-password-btn">
                <FormattedMessage id="user.profile.change_password" />
            </Link>
        </section>
    </div>
}