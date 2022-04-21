import { useFormState } from "react-final-form"
import { useIntl } from "react-intl"
import React from "react"

export type CommonSubmitFieldProps = {
    type?: "button" | "submit" | "reset", // fallbacks to submit,
    label?: string | { id: string },
}

export default (props: CommonSubmitFieldProps) => {
    const intl = useIntl()
    const form = useFormState()

    let { label } = props
    const { type } = props

    let labelRes: string

    if (typeof label === "object")
        labelRes = intl.formatMessage(label)
    else
        labelRes = label ?? ""

    return <div className="any-form__submit">
        <button
            type={type ?? "submit"}
            disabled={form.submitting || form.validating}>
            {labelRes}
        </button>
    </div>
}