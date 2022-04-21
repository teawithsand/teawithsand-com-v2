import React, { useState } from "react"

export default () => {
    const [firstValue, setFirstValue] = useState("")
    const [secondValue, setSecondValue] = useState("")

    return <div className="cs-string-comparator">
        <div className="cs-string-comparator__left">
            <textarea className="cs-string-comparator__input" value={firstValue} onChange={(e) => {
                setFirstValue(e.target.value)
            }}></textarea>
        </div>

        <div className="cs-string-comparator__right">
        <textarea className="cs-string-comparator__input" value={secondValue} onChange={(e) => {
                setSecondValue(e.target.value)
            }}></textarea>
        </div>
        <div className="cs-string-comparator__result">
            {firstValue === secondValue ? <>Strings are equal</> : <>Strings are <b>not</b> equal</>}
        </div>
    </div>
}