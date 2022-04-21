import { Field, useForm } from "react-final-form"
import { FieldArray } from 'react-final-form-arrays'
import React from "react"
import { FormattedMessage, useIntl } from "react-intl"
import classnames from "@app/util/lang/classnames"

export interface WordsFieldProps {
    name: string,
}

export default (props: WordsFieldProps) => {
    const intl = useIntl()
    const { name } = props

    const {
        mutators: { push }
    } = useForm()

    return <div className="any-form__words-input">
        <div className="any-form__words-input__panel">
            <div
                className="any-form__words-input__word-add"
                onClick={() => push(name, undefined)}
            >
                <FormattedMessage id="langka.form.words_field.add_word" />
            </div>
        </div>
        <div className="any-form__words-input__words">
            <FieldArray<string[]> name={name}>
                {
                    ({ fields }) => {
                        if (fields.length === 0) {
                            return <div className="any-form__words-input__no-words">
                                <FormattedMessage id="langka.form.words_field.no_words" />
                            </div>
                        } else {
                            return fields.map((name, index) => {
                                return <div className="any-form__words-input__word" key={name}>
                                    <div className="any-form__words-input__word-input">
                                        <Field<string>
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
                                                        <div className="any-form__input-group">
                                                            <span className="any-form__words-input__badge">#{index + 1}</span>
                                                            <input
                                                                {...input}
                                                                placeholder={intl.formatMessage({
                                                                    id: "langka.form.words.words_field.word.placeholder"
                                                                })}
                                                                type="text"
                                                            />
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
                                    </div>
                                    <div
                                        className="any-form__words-input__word-delete"
                                        onClick={() => fields.remove(index)}
                                    >
                                        <FormattedMessage id="langka.form.words_field.delete_word" />
                                    </div>
                                </div>
                            })
                        }
                    }
                }
            </FieldArray>
        </div>
    </div>
}