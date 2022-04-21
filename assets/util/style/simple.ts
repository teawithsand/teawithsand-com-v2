import classnames from "../lang/classnames"

export const simpleBemParent = (element: string) => {
    return {
        root: () => simpleBemElement(element, ""),
        child: (name: string) => simpleBemElement(element, name)
    }
}

export const simpleBemElement = (element: string, subelement?: string) => {
    if (subelement) {
        element = `${element}__${subelement}`
    }

    return (...modifiers: string[]) => classnames(...[
        element,
        ...(modifiers || []).map(modifier => {
            if (modifier) {
                return `${element}--${modifier}`
            } else {
                return modifier
            }
        })
    ])

}