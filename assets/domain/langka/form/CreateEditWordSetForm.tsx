import CommonSubmitField from "@app/util/form/field/CommonSubmitField"
import CommonTextField from "@app/util/form/field/CommonTextField"
import { processFormData, useFormHelper } from "@app/util/form/fieldProcess"
import { AnyFormProps } from "@app/util/form/form"
import React, { useMemo } from "react"
import { Form } from "react-bootstrap"
import { Form as FinalForm } from "react-final-form"
import { useIntl } from "react-intl"

import livrRules from "@app/exported/livr/langka/word_set_create.json"
import { compileLIVRValidator, mapMessagesValidator, MessagesMap as MessageMap, translateMessageMap, Validator } from "@app/util/validator/typed_validator"
import CommonFormError from "@app/util/form/field/CommonFormError"
import { RemoteWordSetCreateEditData } from "@app/exported/typings/api"

export type EditWordSetData = RemoteWordSetCreateEditData

const formValidator = (): Validator<EditWordSetData> => {
    return compileLIVRValidator<EditWordSetData>(livrRules)
}

const messageMap: MessageMap = {

}

export default (props: AnyFormProps<EditWordSetData>) => {
    const intl = useIntl()
    const helper = useFormHelper(props)

    const validator = useMemo(() =>
        mapMessagesValidator(
            formValidator(),
            {
                map: translateMessageMap(intl, messageMap),
            },
        ), [intl])

    const processData = (data: Partial<EditWordSetData>): EditWordSetData => (processFormData({
        fromLanguage: data.fromLanguage ?? "",
        toLanguage: data.toLanguage ?? "",
        name: data.name ?? "",
        description: data.description ?? "",
        wordTuples: [],
    }))

    return <FinalForm<EditWordSetData>
        initialValues={props.initialData ?? {}}
        onSubmit={data => helper.runOnSubmit(processData(data))}
        validate={(data) => {
            return validator.validate(processData(data))
        }}
        render={({ handleSubmit }) => (<Form
            onSubmit={handleSubmit}
            className="langka-word-set-create-edit-form"
        >
            <CommonFormError />

            <CommonTextField
                name="name"
                placeholder={{ id: "langka.form.wordset.create_edit.name.placeholder" }}
                label={{ id: "langka.form.wordset.create_edit.name.label" }}
                type="text"
            />

            <CommonTextField
                name="fromLanguage"
                placeholder={{ id: "langka.form.wordset.create_edit.from_language.placeholder" }}
                label={{ id: "langka.form.wordset.create_edit.from_language.label" }}
                type="text"
            />

            <CommonTextField
                name="toLanguage"
                placeholder={{ id: "langka.form.wordset.create_edit.to_language.placeholder" }}
                label={{ id: "langka.form.wordset.create_edit.to_language.label" }}
                type="text"
            />

            <CommonTextField
                name="description"
                placeholder={{ id: "langka.form.wordset.create_edit.description.placeholder" }}
                label={{ id: "langka.form.wordset.create_edit.description.label" }}
                type="textarea"
            />

            <CommonSubmitField
                label={{ id: "langka.form.wordset.create_edit.submit.label" }}
            />
        </Form>)}
    />
}