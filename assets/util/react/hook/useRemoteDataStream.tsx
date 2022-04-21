import { useEffect, useState } from "react"

export type RemoteDataStreamLoadInfo<PI> = {
    type: "idle",
    pageInfo?: undefined,
    error?: undefined,
} | {
    type: "pending",
    pageInfo: PI,
    error?: undefined,
} | {
    type: "ok",
    pageInfo: PI,
    error?: undefined,
} | {
    type: "error",
    pageInfo: PI,
    error?: undefined,
}

export type RemoteDataStreamOperations<S, PI> = {
    // Like set state, but uses initializer.
    resetState: () => void,
    setState: (state: S) => void,
    enqueueLoad: (pageInfo: PI) => void,
}

export type RemoteDataStream<S, PI> = [
    S,
    RemoteDataStreamLoadInfo<PI>,
    RemoteDataStreamOperations<S, PI>
]

export interface RemoteDataStreamOptions<S, T, PI, PA> {
    // note: there is no guarantee that this function will be called only once
    // in general it should be: dependent on params only and thus always return same result for given set of parameters
    stateFactory: (params: PA) => S,

    // note: this function will be called only once during component initialization
    initialParametersFactory?: (parameters: PA) => PI,
    loader: (pageInfo: PI, params: PA) => Promise<T>,

    foldValue: (prevState: S, data: T) => S,
    foldError?: (prevState: S, error: any) => S,
}

/**
 * Loaded, which loads multiple smaller resources and folds them into single bigger one.
 * It resets it's state, when one of parameters is changed.
 */
export const useRemoteDataStream = <S, T, PI, PA extends Array<any>>(
    options: RemoteDataStreamOptions<S, T, PI, PA>,
    parameters: PA,
): RemoteDataStream<S, PI> => {
    const [isStateInitial, setIsStateInitial] = useState(true)
    const [state, setState] = useState<S>(() => options.stateFactory(parameters))
    const [loadQueue, setLoadQueue] = useState<PI[]>([])
    const [lastLoadInfo, setLastLoadInfo] = useState<RemoteDataStreamLoadInfo<PI>>({
        type: "idle",
    })

    const getQueueElementAndLoad = async (isClosed: { current: boolean }) => {
        if (isClosed.current)
            return

        if (loadQueue.length === 0)
            return
        const pageInfo = loadQueue[0]
        setLastLoadInfo({
            type: "pending",
            pageInfo,
        })

        try {
            try {
                const v = await options.loader(pageInfo, parameters)
                if (!isClosed.current) {
                    setState(options.foldValue(state, v))
                    setIsStateInitial(false)
                    
                    setLastLoadInfo({
                        type: "ok",
                        pageInfo,
                    })
                }
            } catch (e) {
                if (!isClosed.current) {
                    const folder = options.foldError ?? ((state, _) => state)
                    setState(folder(state, e))
                    setIsStateInitial(false)

                    setLastLoadInfo({
                        type: "error",
                        pageInfo,
                        error: e,
                    })
                }
            }
        } finally {
            if (!isClosed.current) {
                const newQueue = [...loadQueue]
                newQueue.shift()
                setLoadQueue(newQueue)
            }
        }
    }

    const opResetState = () => {
        setState(options.stateFactory(parameters))
        setIsStateInitial(true)
    }
    const opSetState = (state: S) => {
        setState(state)
        setIsStateInitial(false)
    }
    const opEnqueueLoad = (pageInfo: PI) => {
        setLoadQueue([...loadQueue, pageInfo])
    }


    useEffect(() => {
        if (!isStateInitial) {
            setState(options.stateFactory(parameters))
            setIsStateInitial(true)
        }
        setLoadQueue([])
        setLastLoadInfo({
            type: "idle",
        })

        if(options.initialParametersFactory) {
            opEnqueueLoad(options.initialParametersFactory(parameters))
        }

        return () => { }
    }, parameters)

    useEffect(() => {
        let isClosed = { current: false }

        if (loadQueue.length > 0) {
            getQueueElementAndLoad(isClosed)
        }

        return () => {
            isClosed.current = true
        }
    }, [loadQueue])


    return [
        state,
        lastLoadInfo,
        {
            enqueueLoad: opEnqueueLoad,
            resetState: opResetState,
            setState: opSetState,
        }
    ]
}