import React from "react"
import { FormattedMessage } from "react-intl"

import { useErrorExplainer } from "@app/util/explain/error"

export default (props: {
    error: any,
}) => {
    const { error } = props
    const explainer = useErrorExplainer()
    const explained = explainer(error)

    return <div className="app-error">
        <h1 className="app-error__title">
            <FormattedMessage id="common.loading_error.title" />
        </h1>
        <p className="app-error__message">
            <FormattedMessage id={explained.messageKey} values={explained.messageParams ?? {}} />
        </p>
    </div>
}