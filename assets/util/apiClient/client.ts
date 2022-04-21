import { ApiClientError, responseSuccessOrThrow } from "./error"
import { AuthHeaderName } from "@app/exported/typings/api"
import { ErrorResponse } from "../explain/error"
import { useSelector } from "react-redux"
import { userDataSelector } from "@app/domain/common/redux/user"

export interface ApiRequest<RQ> {
    method: string,
    url: string,
    body: RQ,
    headers?: {
        [key: string]: string
    },
    searchParams?: {
        [key: string]: string,
    }
}

export interface ApiResponse<RS> {
    headers: Headers,
    status: number,
    ok: boolean,
    data: RS,
}

export type FallibleApiResponse<RS> = ApiResponse<RS | ErrorResponse>

export interface FetcherOptions {
    url: string,
    method?: string,
}

export interface SimpleHttpClient {
    apiFetch: <RQ, RS>(req: ApiRequest<RQ>) => Promise<ApiResponse<RS>>,

    apiFetcher: <RQ, RS>(options: FetcherOptions) => (req: RQ) => Promise<ApiResponse<RS>>,
    apiSearchFetcher: <RQ, RS>(options: FetcherOptions) => (req: RQ) => Promise<ApiResponse<RS>>,
}

export const makeApiClient = (token?: string): SimpleHttpClient => {
    const tokenHeaders: any = {}
    if (token) {
        tokenHeaders[AuthHeaderName] = "Bearer " + token
    }
    const apiFetch = async  <RQ, RS>(req: ApiRequest<RQ>): Promise<ApiResponse<RS>> => {
        // TODO(teawithsand): allow remote urls(?)
        const url = new URL(req.url, window.location.origin)
        
        Object.entries(req.searchParams ?? {}).forEach(([k, v]) => {
            url.searchParams.set(k, v)
        })

        try {
            const res = await fetch(url.toString(), {
                method: req.method,
                body: JSON.stringify(req.body),
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json",
                    ...tokenHeaders,
                    ...(req.headers ?? {}),
                }
            })

            const body = await (res.json()
                .catch(e => {
                    throw new ApiClientError("Filed to parse JSON response");
                }))

            return {
                headers: res.headers,
                status: res.status,
                ok: res.ok,
                data: body as RS
            }
        } catch (e) {
            if (e instanceof ApiClientError) {
                throw e
            }

            if (typeof e === "object" && e !== null) {
                throw new ApiClientError(e.message ?? "Unknown error");
            } else {
                throw new ApiClientError(String(e));
            }
        }
    }


    const apiFetcher = <RQ, RS>(options: FetcherOptions): ((req: RQ) => Promise<ApiResponse<RS>>) =>
        async (req: RQ) => {
            const res = await apiFetch<RQ, RS | ErrorResponse>({
                body: req,
                method: options.method || "POST",
                url: options.url,
            })

            return responseSuccessOrThrow(res)
        }

    const apiSearchFetcher = <RQ, RS>(options: FetcherOptions): ((req: RQ) => Promise<ApiResponse<RS>>) =>
        async (req: RQ) => {
            const res = await apiFetch<RQ, RS | ErrorResponse>({
                method: options.method || "GET",
                url: options.url,

                body: undefined,
                searchParams: req as unknown as any,
            })

            return responseSuccessOrThrow(res)
        }


    return {
        apiFetch,
        apiFetcher,
        apiSearchFetcher,
    }
}

export const useApiClient = () => {
    const userData = useSelector(userDataSelector)
    return makeApiClient((userData ?? {}).token ?? "")
}