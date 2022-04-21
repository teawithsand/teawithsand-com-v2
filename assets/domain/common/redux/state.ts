import { RulesAction } from "./rules"
import { UserAction, UserData } from "./user"

export type Action = UserAction | RulesAction

export type State = {
    version: 1,
    userData: UserData | null,

    isAcceptedCookies: boolean,
}

export const initialState: State = {
    version: 1,
    userData: null,
    isAcceptedCookies: false,
}