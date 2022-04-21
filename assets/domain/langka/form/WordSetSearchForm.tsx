import CommonSubmitField from "@app/util/form/field/CommonSubmitField"
import CommonTextField from "@app/util/form/field/CommonTextField"
import { processFormData, useFormHelper } from "@app/util/form/fieldProcess"
import { AnyFormProps } from "@app/util/form/form"
import React from "react"
import { Form } from "react-bootstrap"
import { Form as FinalForm } from "react-final-form"
import CommonFormError from "@app/util/form/field/CommonFormError"

export type SearchWordSetData = {
    ownerName: string,
    wordSetName: string,
}

export default (props: AnyFormProps<SearchWordSetData>) => {
    const helper = useFormHelper(props)

    const processData = (data: Partial<SearchWordSetData>): SearchWordSetData => (processFormData({
        ownerName: data.ownerName ?? "",
        wordSetName: data.wordSetName ?? "",
    }))

    return <FinalForm<SearchWordSetData>
        initialValues={props.initialData ?? {}}
        onSubmit={data => helper.runOnSubmit(processData(data))}
        render={({ handleSubmit }) => (<Form
            onSubmit={handleSubmit}
            className="langka-word-set-search-form"
        >
            <CommonFormError />

            <CommonTextField
                name="ownerName"
                placeholder={{ id: "langka.form.wordset.search.owner_name.placeholder" }}
                label={{ id: "langka.form.wordset.search.owner_name.label" }}
                type="text"
            />

            <CommonTextField
                name="wordSetName"
                placeholder={{ id: "langka.form.wordset.search.word_set_name.placeholder" }}
                label={{ id: "langka.form.wordset.search.word_set_name.label" }}
                type="text"
            />

            <CommonSubmitField
                label={{ id: "langka.form.wordset.search.submit.label" }}
            />
        </Form>)}
    />
}