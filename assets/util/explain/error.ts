import { IntlShape, useIntl } from "react-intl"
import { ApiClientError, ErrorResponseApiClientError, InvalidStatusApiClientError, isErrorResponse } from "../apiClient/error"

export interface ErrorResponse {
    debugMessage: string,
    messageKey: string,
    messageParams: { [key: string]: string } | null,
}

export interface ExplainedError {
    debugMessage: string,
    messageKey: string,
    messageParams: { [key: string]: string } | null,

    error: any,
}

export interface TranslatedExplainedError extends ExplainedError {
    translatedMessage: string,
}

const statusMessageMap: { [key: number]: string } & { _default: string } = {
    404: "common.error.http.not_found",
    500: "common.error.http.internal_server_error",
    414: "common.error.http.unprocessable_entity",
    403: "common.error.http.forbidden",
    401: "common.error.http.unauthorized",
    _default: "common.error.http.unknown",
}

export const explainError = (e: any): ExplainedError => {
    if (e instanceof ErrorResponseApiClientError) {
        const { response: { data } } = e

        return {
            debugMessage: "ErrorResponseApiClientError: " + data.debugMessage,
            messageKey: data.messageKey,
            messageParams: data.messageParams || {},

            error: e,
        }
    } else if (e instanceof InvalidStatusApiClientError) {
        const { response } = e
        const { data } = response

        if (isErrorResponse(data)) {
            return {
                debugMessage: "InvalidStatusApiClientError: " + data.debugMessage,
                messageKey: data.messageKey,
                messageParams: data.messageParams || {},

                error: e,
            }
        } else {
            return {
                debugMessage: `InvalidStatusApiClientError: ${response.status}`,
                messageKey: statusMessageMap[response.status] ?? statusMessageMap._default,
                messageParams: {
                    status: `${response.status}`,
                },
                error: e,
            }
        }
    } else if (e instanceof ApiClientError) {
        return {
            debugMessage: "ApiClientError: " + e.message,
            messageKey: "common.error.front.api_error",
            messageParams: {},

            error: e,
        }
    } else {
        console.error("Cannot explain", e)
        return {
            debugMessage: "Unknown error",
            messageKey: "common.error.front.unknown",
            messageParams: {},

            error: e,
        }
    }
}

export const translateError = (intl: IntlShape, e: ExplainedError): TranslatedExplainedError => {
    return {
        ...e,
        translatedMessage: intl.formatMessage({
            id: e.messageKey,
        }, {
            ...(e.messageParams || {})
        })
    }
}

export const useErrorExplainer = (): (e: any) => TranslatedExplainedError => {
    const intl = useIntl()
    return (e) => translateError(intl, explainError(e))
}