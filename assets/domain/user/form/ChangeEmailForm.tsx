import CommonSubmitField from "@app/util/form/field/CommonSubmitField"
import CommonTextField from "@app/util/form/field/CommonTextField"
import { processFormData, useFormHelper } from "@app/util/form/fieldProcess"
import { AnyFormProps } from "@app/util/form/form"
import React, { useMemo } from "react"
import { Form } from "react-bootstrap"
import { Form as FinalForm } from "react-final-form"
import { useIntl } from "react-intl"

import registerRules from "@app/exported/livr/user/change_email.json"
import { compileLIVRValidator, mapMessagesValidator, MessagesMap as MessageMap, translateMessageMap, Validator } from "@app/util/validator/typed_validator"
import CommonFormError from "@app/util/form/field/CommonFormError"
import RecaptchaField from "@app/util/form/field/RecaptchaField"


export type ChangeEmailFormData = {
    email: string,
    captchaResponse: string,
};

const registerValidator = (): Validator<ChangeEmailFormData> => {
    return compileLIVRValidator<ChangeEmailFormData>(registerRules)
}

const messageMap: MessageMap = {
    email: {
        _default: "user.validator.email.invalid",
    },
}

export default (props: AnyFormProps<ChangeEmailFormData>) => {
    const intl = useIntl()
    const helper = useFormHelper(props)

    const validator = useMemo(() =>
        mapMessagesValidator(
            registerValidator(),
            {
                map: translateMessageMap(intl, messageMap),
            },
        ), [intl])

    return <FinalForm<ChangeEmailFormData>
        onSubmit={data => helper.runOnSubmit({
            email: data.email ?? "",
            captchaResponse: data.captchaResponse ?? "",
        })}

        validate={(data) => validator.validate(processFormData(data))}
        render={({ handleSubmit }) => (<Form
            onSubmit={handleSubmit}
            className="user-change-email-form"
        >
            <CommonFormError />

            <CommonTextField
                name="email"
                placeholder={{ id: "user.form.change_email.email.placeholder" }}
                label={{ id: "user.form.change_email.email.label" }}
                type="email"
            />

            <RecaptchaField 
                name="captchaResponse"
            />
            <CommonSubmitField
                label={{ id: "user.form.change_email.submit.label" }}
            />
        </Form>)}
    />
}