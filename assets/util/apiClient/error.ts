import { ErrorResponse } from "../explain/error";
import BaseError from "../lang/error";
import { ApiResponse } from "./client";

export class ApiClientError extends BaseError {
    constructor(message: string) {
        super(message);
    }
}

export class ErrorResponseApiClientError extends ApiClientError {
    constructor(message: string, public readonly response: ApiResponse<ErrorResponse>) {
        super(message);
    }
}

export class InvalidStatusApiClientError extends ApiClientError {
    constructor(message: string, public readonly response: ApiResponse<any>) {
        super(message);
    }
}

export const responseStatusOkOrThrow = <R>(response: ApiResponse<R>): ApiResponse<R> => {
    if (Math.floor(response.status / 100) !== 2) {
        throw new InvalidStatusApiClientError(`InvalidStatusApiClientError: ${response.status}`, response)
    }
    return response as ApiResponse<R>
}

export const isErrorResponse = (res: any) => {
    if (typeof res !== "object" || res === null)
        return false;

    return typeof res.debugMessage === "string" &&
        typeof res.messageKey === "string"
}

export const responseOkOrThrow = <R>(response: ApiResponse<R | ErrorResponse>): ApiResponse<R> => {
    if (isErrorResponse(response.data)) {
        throw new ErrorResponseApiClientError("ErrorResponseApiClientError", response as ApiResponse<ErrorResponse>)
    }

    return response as ApiResponse<R>
}


export const responseSuccessOrThrow = <R>(response: ApiResponse<R | ErrorResponse>): ApiResponse<R> => {
    return responseStatusOkOrThrow(responseOkOrThrow(response))
}