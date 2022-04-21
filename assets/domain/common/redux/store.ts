import { useDispatch } from "react-redux"
import { applyMiddleware, createStore, Store } from "redux"

import { persistStore, persistReducer, Persistor } from "redux-persist";
import storage from "redux-persist/lib/storage";

import {
    createStateSyncMiddleware,
} from "redux-state-sync";
import { PersistPartial } from "redux-persist/es/persistReducer"
import { initialState, State } from "./state";
import rootReducer from "./rootReducer";

export const makeDefaultState = (): State => {
    return initialState
}

let wasCalled = false
const configureStore = (): {
    persistor: Persistor,
    store: Store,
} => {
    if (wasCalled)
        throw new Error("this fn can be called only once");
    wasCalled = true

    const initialState = makeDefaultState()

    const reducer = persistReducer({
        key: "root",
        storage,
        blacklist: ["toasts"],
        throttle: 3000,
    }, rootReducer)

    const store = createStore(
        reducer,
        initialState as (State & PersistPartial),
        applyMiddleware(createStateSyncMiddleware({
            predicate: (action) => {
                if (/^persist\//i.test(action.type))
                    return false;

                if (/^TWSAPI\/toast/i.test(action.type))
                    return false;

                return true;
            }
        }))
    )

    const persistor = persistStore(store)

    return { persistor, store: store }
}

// TODO(teawithsand): make this typed
export const useStoreDispatch =
    () => useDispatch()

const res = configureStore()

export default () => res

const { store, persistor } = res
export { store, persistor } 