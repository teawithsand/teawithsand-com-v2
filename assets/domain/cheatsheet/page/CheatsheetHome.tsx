import React, { useState } from "react"
import { Collapse } from "react-bootstrap"
import { FormattedMessage } from "react-intl"
import StringComparator from "../component/StringComparator"

export default () => {
    const [showStringComparator, setShowStringComparator] = useState(false)

    return <div className="page-cheatsheet-home">
        <header className="page-cheatsheet-home__header">
            <h1>
                <FormattedMessage id="cheatsheet.page.home.title" />
            </h1>
            <p>
                <FormattedMessage id="cheatsheet.page.home.subtitle" />
            </p>
        </header>
        <main className="page-cheatsheet-home__features">
            <div className="page-cheatsheet-home__features__feature">
                <button onClick={() => {
                    setShowStringComparator(!showStringComparator)
                }} className="page-cheatsheet-home__features__btn">
                    <FormattedMessage id="cheatsheet.page.home.string_comparator.toggle" />
                </button>
                <Collapse in={showStringComparator}>
                    <div className="page-cheatsheet-home__features__dropdown">
                        <StringComparator />
                    </div>
                </Collapse>
            </div>
        </main>
    </div>
}