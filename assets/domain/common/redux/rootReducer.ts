import { Action, State } from "./state";

export default (state: State, action: Action): State => {
    switch (action.type) {
        case "TWSAPI/user/set-user-data":
            return {
                ...state,
                userData: action.data,
            }
        case "TWSAPI/user/unset-user-data":
            return {
                ...state,
                userData: null,
            }
        case "TWSAPI/rules/accept-cookies":
            return {
                ...state,
                isAcceptedCookies: true,
            }
        default:
            return state
    }
}