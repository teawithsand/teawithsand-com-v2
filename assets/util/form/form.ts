import { FORM_ERROR } from "final-form";

export interface AnyFormResult {
    [FORM_ERROR]?: string,
}

export interface AnyFormProps<T> {
    initialData?: T,
    onSubmit: (data: T) => Promise<any>,
    onSuccess?: (res: any) => Promise<void> | void,
}