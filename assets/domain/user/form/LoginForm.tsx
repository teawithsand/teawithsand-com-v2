import CommonFormError from "@app/util/form/field/CommonFormError"
import CommonSubmitField from "@app/util/form/field/CommonSubmitField"
import CommonTextField from "@app/util/form/field/CommonTextField"
import { useFormHelper } from "@app/util/form/fieldProcess"
import { AnyFormProps } from "@app/util/form/form"
import React from "react"
import { Form } from "react-bootstrap"
import { Form as FinalForm } from "react-final-form"

export interface LoginFormData {
    username: string,
    password: string,
}

export default (props: AnyFormProps<LoginFormData>) => {
    const helper = useFormHelper(props)

    return <FinalForm<LoginFormData>
        onSubmit={data => helper.runOnSubmit({
            username: data.username ?? "",
            password: data.password ?? "",
        })}
        
        render={({ handleSubmit }) => (<Form
            onSubmit={handleSubmit}
            className="user-login-form"
        >
            <CommonFormError />

            <CommonTextField
                name="username"
                placeholder={{ id: "user.form.login.username.placeholder" }}
                label={{ id: "user.form.login.username.label" }}
            />


            <CommonTextField
                name="password"
                placeholder={{ id: "user.form.login.password.placeholder" }}
                label={{ id: "user.form.login.password.label" }}
                type="password"
            />

            <CommonSubmitField
                label={{ id: "user.form.login.submit.label" }}
            />
        </Form>)}
    />
}