import { getEndpointPath, userRegisterEndpoint, langkaListOwnedWordSetsEndpoint, langkaListPublicWordSetsEndpoint, langkaCreateWordSetEndpoint } from "@app/domain/common/endpoints"
import { userDataSelector } from "@app/domain/common/redux/user"
import React from "react"
import { FormattedMessage } from "react-intl"
import { useSelector } from "react-redux"
import { Link } from "react-router-dom"

export default () => {
    const userData = useSelector(userDataSelector)

    const langkaJoin = <div className="page-langka-home__features__feature">
        <h3>
            <FormattedMessage id="langka.page.home.join.title" />
        </h3>
        <p>
            <FormattedMessage id="langka.page.home.join.description" />
        </p>
        <Link to={getEndpointPath(userRegisterEndpoint, null)} className="feature__btn">
            <FormattedMessage id="langka.page.home.join.button" />
        </Link>
    </div>

    const langkaJoined = <>
        <div className="page-langka-home__features__feature">
            <h3>
                <FormattedMessage id="langka.page.home.browse_owned.title" />
            </h3>
            <p>
                <FormattedMessage id="langka.page.home.browse_owned.description" />
            </p>
            <Link to={getEndpointPath(langkaListOwnedWordSetsEndpoint, null)} className="feature__btn">
                <FormattedMessage id="langka.page.home.browse_owned.button" />
            </Link>
        </div>

        <div className="page-langka-home__features__feature">
            <h3>
                <FormattedMessage id="langka.page.home.create.title" />
            </h3>
            <p>
                <FormattedMessage id="langka.page.home.create.description" />
            </p>
            <Link to={getEndpointPath(langkaCreateWordSetEndpoint, null)} className="feature__btn">
                <FormattedMessage id="langka.page.home.create.button" />
            </Link>
        </div>
    </>

    return <div className="page-langka-home">
        <header className="page-langka-home__header">
            <h1>
                <FormattedMessage id="langka.page.home.title" />
            </h1>
            <p>
                <FormattedMessage id="langka.page.home.subtitle" />
            </p>
        </header>
        <main className="page-langka-home__features">
            <div className="page-langka-home__features__feature">
                <h3>
                    <FormattedMessage id="langka.page.home.browse.title" />
                </h3>
                <p>
                    <FormattedMessage id="langka.page.home.browse.description" />
                </p>
                <Link to={getEndpointPath(langkaListPublicWordSetsEndpoint, null)} className="feature__btn">
                    <FormattedMessage id="langka.page.home.browse.button" />
                </Link>
            </div>
            {userData ? langkaJoined : langkaJoin}
        </main>
    </div>
}