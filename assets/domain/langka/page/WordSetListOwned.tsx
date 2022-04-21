import ExplainedLoadingError from "@app/domain/common/component/ExplainedLoadingError"
import LoadingError from "@app/domain/common/component/LoadingError"
import LoadingSpinner from "@app/domain/common/component/LoadingSpinner"
import { userDataSelector } from "@app/domain/common/redux/user"
import { userMustLogInError } from "@app/domain/user/api/error"
import { RemoteWordSetSecretSummaryProjection } from "@app/exported/typings/api"
import { useApiClient } from "@app/util/apiClient/client"
import { useRemoteDataStream } from "@app/util/react/hook/useRemoteDataStream"
import { simpleBemParent } from "@app/util/style/simple"
import React from "react"
import { FormattedMessage } from "react-intl"
import { useSelector } from "react-redux"
import { useSearchParams } from "react-router-dom"
import LangkaWordSetApiClient from "../api/LangkaWordSetApiClient"
import WordSetCard from "../component/WordSetCard"
import WordSetSearchForm from "../form/WordSetSearchForm"


const bemParent = simpleBemParent("page-langka-word-set-owned-list")
const bemRoot = bemParent.root()
const bemHeader = bemParent.child("header")
const bemFooter = bemParent.child("footer")
const bemEntries = bemParent.child("entries")
const bemSearch = bemParent.child("search")

export default () => {
    const [paramters, setParameters] = useSearchParams({

    })

    const userData = useSelector(userDataSelector)

    const page = (parseInt(paramters.get("page")) || 0)
    const perPage = 30

    const typedParameters = {
        ownerName: paramters.get("ownerName") ?? "",
        wordSetName: paramters.get("wordSetName") ?? "",
        destinationLanguage: paramters.get("destinationLanguage") ?? "",
        sourceLanguage: paramters.get("sourceLanguage") ?? "",
    }

    const client = useApiClient()
    const wordSetClient = new LangkaWordSetApiClient(client)

    const [state, info, ops] = useRemoteDataStream({
        stateFactory: (): {
            isAfterLast: boolean,
            entries: RemoteWordSetSecretSummaryProjection[],
        } => ({
            isAfterLast: true,
            entries: [],
        }),
        initialParametersFactory: () => page,

        loader: async (page: number): Promise<RemoteWordSetSecretSummaryProjection[]> => {
            if (!userData)
                return []
            const res = await wordSetClient.getOwnedWordSetList({
                ...typedParameters,
                offset: page * perPage,
                limit: perPage
            })

            return res.data.entries ?? []
        },
        foldValue: (state, data) => {
            return {
                isAfterLast: data.length === 0,
                entries: [...state.entries, ...data]
            }
        }
    }, Object.values(typedParameters))

    if (!userData) {
        return <ExplainedLoadingError error={userMustLogInError} />
    }

    let footerContent = null
    if (info.type === "error") {
        footerContent = <LoadingError error={info.error} />
    } else if (info.type === "pending") {
        footerContent = <LoadingSpinner />
    } else if (state.isAfterLast) {
        footerContent = <FormattedMessage id="langka.page.wordset.list_owned.no_more_entries" />
    } else {
        footerContent = <button className="btn btn-primary" onClick={() => ops.enqueueLoad(page + 1)}>
            <FormattedMessage id="langka.page.wordset.list_owned.load_more" />
        </button>
    }

    return <div className={bemRoot()}>
        <header className={bemHeader()}>
            <h1>
                <FormattedMessage id="langka.page.wordset.list_owned.title" />
            </h1>
        </header>
        <section className={bemSearch()}>
            <WordSetSearchForm
                initialData={{
                    ownerName: typedParameters.ownerName,
                    wordSetName: typedParameters.wordSetName,
                }}
                onSubmit={async (data) => {
                    setParameters({
                        ...typedParameters,
                        ownerName: data.ownerName,
                        wordSetName: data.wordSetName,
                    })
                }}
            />
        </section>
        <main className={bemEntries()}>
            {
                state.entries.length > 0 ?
                    state.entries.map((e) =>  <WordSetCard data={e} type="secret" key={e.id} 
                        showShowButton={true}
                        showEditButton={e.owner.id === userData?.id}
                    />)
                    :
                    <div>
                        <FormattedMessage id="langka.page.wordset.list_owned.no_entries" />
                    </div>
            }
        </main>
        <section className={bemFooter()}>
            {footerContent}
        </section>
    </div>
}