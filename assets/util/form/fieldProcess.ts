import { FORM_ERROR } from "final-form"
import { useErrorExplainer } from "../explain/error"
import { AnyFormProps, AnyFormResult } from "./form"

export const processFieldValue = (v: string): string => {
    return v.trim()
}
export const processFormData = <T extends object>(v: T): T => {
    const res: T = { ...v }
    for (const k in v) {
        if (typeof v[k] === "string") {
            (res as unknown)[k] = processFieldValue(v[k] as unknown as string)
        }
    }
    return res
}

export interface FormHelper<T> {
    runOnSubmit: (data: T) => Promise<AnyFormResult>
}

export const useFormHelper = <T>(props: AnyFormProps<T>): FormHelper<T> => {
    const explainer = useErrorExplainer()

    return {
        runOnSubmit: async (data) => {
            try {
                const res = await props.onSubmit(data)
                if (props.onSuccess) {
                    await (props.onSuccess(res) || Promise.resolve());
                }
            } catch (e) {
                const explained = explainer(e)
                return {
                    [FORM_ERROR]: explained.translatedMessage,
                }
            }

            return {}
        }
    }
}