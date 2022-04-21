import { Field } from "react-final-form"
import { FormattedMessage, useIntl } from "react-intl"
import React from "react"
import classnames from "@app/util/lang/classnames"

export type CommonTextFieldProps = {
    name: string,
    type?: string, // fallbacks to text,
    placeholder?: string | { id: string },
    label?: string | { id: string },
}

export default (props: CommonTextFieldProps) => {
    const intl = useIntl()

    let { placeholder, label } = props
    const { type, name } = props

    let placeholderRes: string
    let labelRes: string
    if (typeof placeholder === "object")
        placeholderRes = intl.formatMessage(placeholder)
    else
        placeholderRes = placeholder ?? ""

    if (typeof label === "object")
        labelRes = intl.formatMessage(label)
    else
        labelRes = label ?? ""

    return <Field<string>
        name={name}
    >
        {
            ({ input, meta }) => (
                <div className={
                    classnames(
                        "any-form__input",
                        meta.touched ? "any-form__input--touched" : null,
                        meta.error ? "any-form__input--error" : null,
                    )
                }>
                    <label>
                        {labelRes}
                    </label>
                    <div className="any-form__input-group">
                        {type === "textarea" ? <textarea
                            {...input}
                            placeholder={placeholderRes}></textarea>
                            :
                            <input
                                {...input}
                                placeholder={placeholderRes}
                                type={type ?? "text"}
                            />}
                    </div>
                    {meta.error && <span
                        data-touched={meta.touched}
                        className="any-form__validation-error">
                        {meta.error}
                    </span>}
                </div>
            )
        }
    </Field>
}