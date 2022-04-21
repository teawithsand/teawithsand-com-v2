import { useFormState } from "react-final-form"
import { useIntl } from "react-intl"
import React from "react"

export type CommonFormErrorProps = {
    title?: string | { id: string },
}

export default (props: CommonFormErrorProps) => {
    const intl = useIntl()
    const form = useFormState()

    let { title } = props

    let titleRes: string

    if (typeof title === "object")
        titleRes = intl.formatMessage(title)
    else
        titleRes = title ?? intl.formatMessage({ id: "common.form.submiterror.title" })

    if (form.submitError) {
        return <div className="any-form__top-error">
            <h3>{titleRes}</h3>
            {form.submitError ? form.submitError : null}
        </div>
    } else {
        return <></>
    }
}