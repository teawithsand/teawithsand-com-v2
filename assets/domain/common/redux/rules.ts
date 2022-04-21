import { State } from "./state"

export type RulesAcceptCookiesAction = {
    type: "TWSAPI/rules/accept-cookies",
}

export type RulesAction = RulesAcceptCookiesAction

export const rulesAcceptCookiesAction = (): RulesAcceptCookiesAction => ({
    type: "TWSAPI/rules/accept-cookies",
})

export const rulesIsAcceptedCookiesSelector = (s: State) => s.isAcceptedCookies