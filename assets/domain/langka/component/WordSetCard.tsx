import { langkaEditWordSetEndpoint, langkaShowPublicWordSetEndpoint, langkaShowSecretWordSetEndpoint, useEndpointNavigate } from "@app/domain/common/endpoints"
import { RemoteWordSetPublicSummaryProjection, RemoteWordSetSecretSummaryProjection } from "@app/exported/typings/api"
import { defaultDateTimeFormatProps } from "@app/util/i18n/options"
import { simpleBemParent } from "@app/util/style/simple"
import React from "react"
import { FormattedMessage, useIntl } from "react-intl"

const bemParent = simpleBemParent("c-langka-word-set-card")
const bemRoot = bemParent.root()
const bemHeader = bemParent.child("header")
const bemMeta = bemParent.child("meta")
const bemControls = bemParent.child("controls")

const buttonBemParent = simpleBemParent("button")
const buttonBemRoot = buttonBemParent.root()

interface WordSetCardCommonProps {
    showShowButton?: boolean,
    showEditButton?: boolean,
}

export default (props: ({
    type: "public",
    data: RemoteWordSetPublicSummaryProjection,
} | {
    type: "secret",
    data: RemoteWordSetSecretSummaryProjection,
}) & WordSetCardCommonProps) => {
    const { data, type, showEditButton, showShowButton } = props
    const navigate = useEndpointNavigate()
    const intl = useIntl()

    return <div className={bemRoot()}>
        <div className={bemHeader()}>
            <h4>{data.name}</h4>
        </div>
        <div className={bemMeta()}>
            <div>
                <FormattedMessage id="component.langka.word_set_card.created_at" values={{
                    createdAt: intl.formatDate(new Date(data.createdAt), {
                        ...defaultDateTimeFormatProps,
                    })
                }} />
            </div>
            <div>
                <FormattedMessage id="component.langka.word_set_card.from_language" values={{
                    fromLanguage: data.fromLanguage
                }} />
            </div>
            <div>
                <FormattedMessage id="component.langka.word_set_card.to_language" values={{
                    toLanguage: data.toLanguage
                }} />
            </div>
            <div>
                <FormattedMessage id="component.langka.word_set_card.owned_by" values={{
                    owner: data.owner.publicName,
                }} />
            </div>
        </div>
        <div className={bemControls()}>
            {showShowButton ? <button className={buttonBemRoot("show")}
                onClick={() => {
                    if (type === "public") {
                        navigate(langkaShowPublicWordSetEndpoint, { id: data.id })
                    } else if (type === "secret") {
                        navigate(langkaShowSecretWordSetEndpoint, { id: data.id })
                    }
                }}
            >
                <FormattedMessage id="component.langka.word_set_card.show" />
            </button> : null}
            {showEditButton ? <button className={buttonBemRoot("edit")}
                onClick={() => {
                    navigate(langkaEditWordSetEndpoint, { id: data.id })
                }}
            >
                <FormattedMessage id="component.langka.word_set_card.edit" />
            </button> : null}
        </div>
    </div>
}