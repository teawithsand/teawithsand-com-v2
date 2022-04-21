import { NavigateOptions, useNavigate } from "react-router"

// TODO(teawithsand): add list support
export type Params = void | Record<string, string>

export type Endpoint<P extends Params = void> = {
    pattern: string,
}

// these have no common prefix
export const commonHomeEndpoint: Endpoint = { pattern: "/" }
export const commonPortfolioHomeEndpoint: Endpoint = { pattern: "/portfolio" }
export const commonAboutMeEndpoint: Endpoint = { pattern: "/about-me" }


export const cheatsheetHomeEndpoint: Endpoint = { pattern: "/cheatsheet" }


export const userLoginEndpoint: Endpoint = { pattern: "/user/login" }
export const userRegisterEndpoint: Endpoint = { pattern: "/user/register" }
export const userConfirmRegistrationEndpoint: Endpoint = { pattern: "/user/register/confirm" }
export const userChangeEmailEndpoint: Endpoint = { pattern: "/user/change-email" }
export const userChangePasswordEndpoint: Endpoint = { pattern: "/user/change-password" }
export const userDeleteAccountEndpoint: Endpoint = { pattern: "/user/profile/delete" }
export const userSecretProfileEndpoint: Endpoint<{ id: string }> = { pattern: "/user/profile/secret" }
export const userPublicProfileEndpoint: Endpoint<{ id: string }> = { pattern: "/user/profile/public" }


export const langkaHomeEndpoint: Endpoint = { pattern: "/langka" }
export const langkaBrowseWordsGameEndpoint: Endpoint = { pattern: "/langka/word-sets/browse-words" }


export const langkaListPublicWordSetsEndpoint: Endpoint = { pattern: "/langka/word-sets/public" }
export const langkaListOwnedWordSetsEndpoint: Endpoint = { pattern: "/langka/word-sets/owned" }
export const langkaShowSecretWordSetEndpoint: Endpoint<{ id: string }> = { pattern: "/langka/word-sets/secret/show" }
export const langkaShowPublicWordSetEndpoint: Endpoint<{ id: string }> = { pattern: "/langka/word-sets/public/show" }

export const langkaCreateWordSetEndpoint: Endpoint = { pattern: "/langka/word-sets/new" }
export const langkaEditWordSetEndpoint: Endpoint<{ id: string }> = { pattern: "/langka/word-sets/edit" }

export const useEndpointNavigate = (): (<P extends Params>(endpoint: Endpoint<P>, param: P, options?: NavigateOptions) => void) => {
    const navigate = useNavigate()

    return (endpoint, param, options) => {
        navigate(getEndpointPath(endpoint, param), options)
    }
}

export const getEndpointPath = <P extends Params>(ep: Endpoint<P>, params: P) => {
    if (typeof params === "object" && params !== null) {
        const typedParams = new URLSearchParams(params as any)
        return ep.pattern + "?" + typedParams.toString()
    }
    return ep.pattern
}

export const getEndpointRoute = (ep: Endpoint<any>): string => {
    return ep.pattern
}