import React from "react"
import { FormattedMessage } from "react-intl"

import { ExplainedError } from "@app/util/explain/error"

export default (props: {
    error: ExplainedError,
}) => {
    const { error } = props

    return <div className="app-error">
        <h1 className="app-error__title">
            <FormattedMessage id="common.loading_error.title" />
        </h1>
        <p className="app-error__message">
            <FormattedMessage id={error.messageKey} values={error.messageParams ?? {}} />
        </p>
    </div>
}