import React from "react"
import { useApiClient } from "@app/util/apiClient/client"
import { langkaShowSecretWordSetEndpoint, useEndpointNavigate } from "@app/domain/common/endpoints"
import { FormattedMessage } from "react-intl"
import CreateWordSetForm from "../form/CreateEditWordSetForm"
import LangkaWordSetApiClient from "../api/LangkaWordSetApiClient"
import { RemoteWordSetCreateResponse } from "@app/exported/typings/api"
import { useRemoteData } from "@app/util/react/hook/useRemoteData"
import { useSearchParams } from "react-router-dom"
import LoadingSpinner from "@app/domain/common/component/LoadingSpinner"
import { useSelector } from "react-redux"
import { userDataSelector } from "@app/domain/common/redux/user"
import ExplainedLoadingError from "@app/domain/common/component/ExplainedLoadingError"
import { userMustLogInError } from "@app/domain/user/api/error"
import LoadingError from "@app/domain/common/component/LoadingError"

export default () => {
    const client = useApiClient()
    const wordSetClient = new LangkaWordSetApiClient(client)
    const navigate = useEndpointNavigate()

    const [paramters] = useSearchParams({
        id: "",
    })
    const id = paramters.get("id")
    const userData = useSelector(userDataSelector)

    const loader = useRemoteData({
        loader: async () => {
            if (!userData) return;
            return await wordSetClient.getSecretWordSet({ id })
        }
    }, [])

    if (!userData) {
        return <ExplainedLoadingError
            error={userMustLogInError}
        />
    }

    if (loader.type === "pending") {
        return <LoadingSpinner />
    } else if (loader.type === "error") {
        return <LoadingError error={loader.error} />
    }

    const data = loader.data.data

    return <div className="page-langka-word-set-edit">
        <header className="page-langka-word-set-edit__header">
            <h1>
                <FormattedMessage id="langka.page.wordset.edit.title" />
            </h1>
        </header>
        <main className="page-langka-word-set-edit__form">
            <CreateWordSetForm
                initialData={{
                    description: data.description,
                    fromLanguage: data.fromLanguage,
                    toLanguage: data.toLanguage,
                    name: data.name,
                    wordTuples: data.wordTuples,
                }}
                onSubmit={async (formData) => {
                    return (await wordSetClient.editWordSet({
                        id,
                        ...formData,
                        wordTuples: data.wordTuples, // just to make sure form didn't modify it
                    })).data
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