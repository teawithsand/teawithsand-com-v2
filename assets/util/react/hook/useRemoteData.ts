import { useEffect, useState } from "react"

export type RemoteData<T> = {
    type: "loaded",
    error?: undefined,

    data: T,
    prevData?: undefined,
} | {
    type: "pending",
    error?: undefined,

    data?: undefined,
    prevData: T | undefined,
} | {
    type: "error",
    error: any,

    data?: undefined,
    prevData?: T | undefined,
}

export interface RemoteDataOptions<T, P> {
    loader: (parameters: P) => Promise<T>,
}

/**
 * Loader, which keeps loaded single resource depending on parameters given.
 */
export const useRemoteData = <T, P extends Array<any>>(
    options: RemoteDataOptions<T, P>,
    parameters: P,
): RemoteData<T> => {
    const [isLoadingData, setIsLoadingData] = useState(true)
    const [currentData, setCurrentData] = useState<T | undefined>(undefined)
    const [error, setError] = useState(undefined)

    useEffect(() => {
        setIsLoadingData(true)
        
        let isClosed = false
        options.loader(parameters)
            .then(v => {
                if (!isClosed) {
                    setCurrentData(v)
                    setError(undefined)
                }
            })
            .catch(e => {
                if (!isClosed)
                    setError(e)
            }).finally(() => {
                if (!isClosed)
                    setIsLoadingData(false)
            })
        return () => {
            isClosed = true
        }
    }, parameters)


    if (isLoadingData) {
        return {
            type: "pending",
            prevData: currentData,
        }
    } else if (error) {
        return {
            type: "error",
            error,
            prevData: currentData,
        }
    } else {
        return {
            type: "loaded",
            data: currentData,
        }
    }

}