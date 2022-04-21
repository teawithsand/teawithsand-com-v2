import * as LIVR from "livr"
export const innerCompileValidator = rules => {
    return new LIVR.Validator(rules)
} 