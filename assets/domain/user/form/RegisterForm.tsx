import CommonSubmitField from "@app/util/form/field/CommonSubmitField"
import CommonTextField from "@app/util/form/field/CommonTextField"
import { processFormData, useFormHelper } from "@app/util/form/fieldProcess"
import { AnyFormProps } from "@app/util/form/form"
import React, { useMemo } from "react"
import { Form } from "react-bootstrap"
import { Form as FinalForm } from "react-final-form"
import { useIntl } from "react-intl"

import registerRules from "@app/exported/livr/user/register.json"
import { compileLIVRValidator, mapMessagesValidator, MessagesMap as MessageMap, translateMessageMap, Validator } from "@app/util/validator/typed_validator"
import CommonFormError from "@app/util/form/field/CommonFormError"
import RecaptchaField from "@app/util/form/field/RecaptchaField"
import { RemoteRegisterUserData } from "@app/exported/typings/api"

export type RegisterFormData = RemoteRegisterUserData;

interface InnerRegisterFormData extends RegisterFormData {
    repeatPassword: string
}

const registerValidator = (): Validator<InnerRegisterFormData> => {
    return compileLIVRValidator<InnerRegisterFormData>(registerRules)
}

const messageMap: MessageMap = {
    username: {
        _default: "user.validator.username.invalid",
    },
    email: {
        _default: "user.validator.email.invalid",
    },
}

export default (props: AnyFormProps<RegisterFormData>) => {
    const intl = useIntl()
    const helper = useFormHelper(props)

    const validator = useMemo(() =>
        mapMessagesValidator(
            registerValidator(),
            { map: translateMessageMap(intl, messageMap), },
        ), [intl])

    return <FinalForm<InnerRegisterFormData>
        onSubmit={data => helper.runOnSubmit({
            username: data.username ?? "",
            email: data.email ?? "",
            captchaResponse: data.captchaResponse ?? "",
        })}
        validate={(data) => validator.validate(processFormData(data))}
        render={({ handleSubmit }) => (<Form
            onSubmit={handleSubmit}
            className="user-register-form"
        >
            <CommonFormError />

            <CommonTextField
                name="username"
                placeholder={{ id: "user.form.register.username.placeholder" }}
                label={{ id: "user.form.register.username.label" }}
            />

            <CommonTextField
                name="email"
                placeholder={{ id: "user.form.register.email.placeholder" }}
                label={{ id: "user.form.register.email.label" }}
                type="email"
            />

            <RecaptchaField
                name="captchaResponse"
            />

            <CommonSubmitField
                label={{ id: "user.form.register.submit.label" }}
            />
        </Form>)}
    />
}