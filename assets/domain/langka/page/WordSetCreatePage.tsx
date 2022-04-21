import React from "react"
import { useApiClient } from "@app/util/apiClient/client"
import { langkaShowSecretWordSetEndpoint, useEndpointNavigate } from "@app/domain/common/endpoints"
import { FormattedMessage } from "react-intl"
import CreateWordSetForm from "../form/CreateEditWordSetForm"
import LangkaWordSetApiClient from "../api/LangkaWordSetApiClient"
import { RemoteWordSetCreateResponse } from "@app/exported/typings/api"
import { useSelector } from "react-redux"
import { userDataSelector } from "@app/domain/common/redux/user"
import ExplainedLoadingError from "@app/domain/common/component/ExplainedLoadingError"
import { userMustLogInError } from "@app/domain/user/api/error"

export default () => {
    const userData = useSelector(userDataSelector)
    const client = useApiClient()
    const wordSetClient = new LangkaWordSetApiClient(client)
    const navigate = useEndpointNavigate()

    if (!userData) {
        return <ExplainedLoadingError
            error={userMustLogInError}
        />
    }

    return <div className="page-langka-word-set-create">
        <header className="page-langka-word-set-create__header">
            <h1>
                <FormattedMessage id="langka.page.wordset.create.title" />
            </h1>
        </header>
        <main className="page-langka-word-set-create__form">
            <CreateWordSetForm
                onSubmit={async (formData) => {
                    return (await wordSetClient.createWordSet(formData)).data
                }}
                onSuccess={(res: RemoteWordSetCreateResponse) => {
                    navigate(langkaShowSecretWordSetEndpoint, {
                        id: res.id,
                    })
                }}
            />
        </main>
    </div>
}