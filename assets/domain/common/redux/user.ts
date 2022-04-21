import { State } from "./state"

export type UserData = {
    id: string,
    publicName: string,
    token: string,
}

export type UserSetDataAction = {
    type: "TWSAPI/user/set-user-data",
    data: UserData,
}

export type UserUnsetDataAction = {
    type: "TWSAPI/user/unset-user-data",
}

export type UserAction = UserSetDataAction | UserUnsetDataAction

export const userSetDataAction = (ud: UserData): UserSetDataAction => ({
    type: "TWSAPI/user/set-user-data",
    data: ud,
})

export const unsetUnsetDataAction = (): UserUnsetDataAction => ({
    type: "TWSAPI/user/unset-user-data",
})


export const userDataSelector = (s: State) => s.userData