import CommonSubmitField from "@app/util/form/field/CommonSubmitField"
import CommonTextField from "@app/util/form/field/CommonTextField"
import { processFormData, useFormHelper } from "@app/util/form/fieldProcess"
import { AnyFormProps } from "@app/util/form/form"
import React, { useMemo } from "react"
import { Form } from "react-bootstrap"
import { Form as FinalForm } from "react-final-form"
import { useIntl } from "react-intl"

import registerRules from "@app/exported/livr/user/change_password.json"
import { compileLIVRValidator, mapMessagesValidator, MessagesMap as MessageMap, translateMessageMap, Validator } from "@app/util/validator/typed_validator"
import CommonFormError from "@app/util/form/field/CommonFormError"
import { RemoteConfirmRegistrationData } from "@app/exported/typings/api"


export type ChangePasswordFormData = {
    password: string,
};

interface InnerChangePasswordFormData extends RemoteConfirmRegistrationData {
    repeatPassword: string
}

const registerValidator = (): Validator<InnerChangePasswordFormData> => {
    return compileLIVRValidator<InnerChangePasswordFormData>(registerRules)
}

const messageMap: MessageMap = {
    password: {
        _default: "user.validator.password.invalid",
    },
    repeatPassword: {
        _default: "user.validator.repeat_password.mismatch",
    },
}

export default (props: AnyFormProps<ChangePasswordFormData>) => {
    const intl = useIntl()
    const helper = useFormHelper(props)

    const validator = useMemo(() =>
        mapMessagesValidator(
            registerValidator(),
            {
                map: translateMessageMap(intl, messageMap),
            },
        ), [intl])

    return <FinalForm<InnerChangePasswordFormData>
        onSubmit={data => helper.runOnSubmit({
            password: data.password ?? "",
        })}

        validate={(data) => {
            const res = validator.validate(processFormData(data))
            let secondError: any = {}
            if ((data.repeatPassword ?? "") !== (data.password ?? "")) {
                secondError = { repeatPassword: messageMap.repeatPassword._default }
            }
            return {
                ...res,
                ...secondError,
            }
        }}
        render={({ handleSubmit }) => (<Form
            onSubmit={handleSubmit}
            className="user-change-password-form"
        >
            <CommonFormError />

            <CommonTextField
                name="password"
                placeholder={{ id: "user.form.change_password.password.placeholder" }}
                label={{ id: "user.form.change_password.password.label" }}
                type="password"
            />


            <CommonTextField
                name="repeatPassword"
                placeholder={{ id: "user.form.change_password.repeat_password.placeholder" }}
                label={{ id: "user.form.change_password.repeat_password.label" }}
                type="password"
            />

            <CommonSubmitField
                label={{ id: "user.form.change_password.submit.label" }}
            />
        </Form>)}
    />
}