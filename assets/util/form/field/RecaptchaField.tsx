import { RecaptchaV2 } from 'react-recaptcha-functional'
import React from "react"
import { Field } from 'react-final-form'
import { useIntl } from 'react-intl'
import classnames from '@app/util/lang/classnames'

export type RecaptchaFieldProps = {
    name: string,
    sitekey?: string,
}


export default (props: RecaptchaFieldProps) => {
    const intl = useIntl()
    return <Field
        name={props.name}

        validate={(value) => {
            if (value === "") {
                return intl.formatMessage({ id: "common.util.recaptcha.required" })
            }
        }}
    >
        {({ input, meta }) => (
            <div className={classnames(
                "any-form__input",
                meta.touched ? "any-form__input--touched" : null,
            )}>
                <RecaptchaV2
                    sitekey="6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI"
                    onTokenChanged={(v) => input.onChange(v ?? "")}
                />
            </div>
        )}
    </Field>
}