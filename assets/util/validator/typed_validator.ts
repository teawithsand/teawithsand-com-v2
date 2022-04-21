import { IntlShape } from "react-intl"
import { innerCompileValidator } from "./validator"

type Error = string | string[]

export type Errors = {
    [key: string]: Error
}

export type InnerMessagesMap = {
    [key: string]: string
}

export type MessagesMap = {
    [key: string]: InnerMessagesMap
}

export type MapMessagesOptions = {
    map: MessagesMap,
    defaultMessage?: string,
}

export const translateMessageMap = (intl: IntlShape, msgMap: MessagesMap): MessagesMap => {
    const res = { ...msgMap }
    for (const ok in msgMap) {
        const ires = { ...msgMap[ok] }
        for (const ik in msgMap[ok]) {
            ires[ik] = intl.formatMessage({ id: ires[ik] })
        }
        res[ok] = ires
    }

    return res
}

export type MsgPicker = (field: string, error: string) => string


export interface Validator<T> {
    validate(data: T): Errors
}

export const compileLIVRValidator = <T>(rules: any): Validator<T> => {
    const compiled = innerCompileValidator(rules)
    return {
        validate: (data: T) => {
            compiled.validate(data)
            const errors = { ...compiled.getErrors() }
            for (const k in errors) {
                if (typeof errors[k] === "string")
                    errors[k] = errors[k].toLowerCase()
            }
            return errors
        }
    }
}

export const mapMessagesValidator = <T>(validator: Validator<T>, options: MapMessagesOptions) => {
    const { map, defaultMessage } = options

    let mapStack: InnerMessagesMap[] = []

    const translateMessage = (fieldValue: string): string => {
        const map = mapStack[mapStack.length - 1] ?? {}
        return map[fieldValue] ?? // message
            map["_default"] ?? // default for field
            defaultMessage ?? // default for form
            fieldValue // inner field value
    }

    const translateErrors = (input: Errors): Errors => {
        return input // TODO(teawithsand): implement support for nested errors, preferrably using paths
    }

    return {
        validate: (data: T) => {
            const res = validator.validate(data) ?? {}
            return translateErrors(res)
        }
    }
}

export const pickMessagesValidator = <T>(validator: Validator<T>, picker: MsgPicker) => {
    return {
        validate: (data: T) => {
            const res = validator.validate(data)
            for (const k in res) {
                const v = res[k]
                if (typeof v === "string") {
                    res[k] = picker(k, v)
                } else {
                    res[k] = v.map((err) => picker(k, err))
                }
            }
            return res
        }
    }
}