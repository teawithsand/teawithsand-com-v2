import React, { useRef } from "react"
import { FormattedMessage, useIntl } from "react-intl"
import { Link } from "react-router-dom"
import { cheatsheetHomeEndpoint, getEndpointPath, langkaHomeEndpoint, commonPortfolioHomeEndpoint } from "../endpoints"
import ImageUtil from "@app/util/react/image/ImageUtil"

import phoneImage from "@app/images/svgrepo/phone.svg"
import emailImage from "@app/images/svgrepo/email.svg"
import TeaAnimation from "./TeaAnimation"


export default () => {
    const intl = useIntl()
    // const headerStyle = getBackgroundImageStyle(starryNight)

    const firstSectionRef = useRef<HTMLElement>()

    return <div className="page-home">
        <header className="page-home__header">
            <div className="page-home__header__background">
                <TeaAnimation />
            </div>

            <div className="page-home__header__overlay app-tea-appear">
                <h1>
                    <FormattedMessage id="common.home.header.title" />
                </h1>
                <p>
                    <FormattedMessage id="common.home.header.paragraph" />
                </p>

                <button className="app-tea-appear-later btn btn-primary" onClick={() => {
                    firstSectionRef.current.scrollIntoView({
                        behavior: "smooth",
                        block: "start",
                    })
                }}>
                    <FormattedMessage id="common.home.header.see_more_button" />
                </button>
            </div>
        </header>
        <section className="page-home__hero-about" ref={firstSectionRef}>
            <h2>
                <FormattedMessage id="common.home.aboutpage.title" />
            </h2>
            <p>
                <FormattedMessage id="common.home.aboutpage.paragraph" />
            </p>
        </section>
        <main className="page-home__hero-features">

            <div className="page-home__hero-features__feature">
                <h3>
                    <FormattedMessage id="common.home.features.langka.title" />
                </h3>
                <p>
                    <FormattedMessage id="common.home.features.langka.paragraph" />
                </p>

                <Link to={getEndpointPath(langkaHomeEndpoint, null)}
                    className="feature__btn">
                    <FormattedMessage id="common.home.features.langka.button" />
                </Link>
            </div>

            <div className="page-home__hero-features__feature">
                <h3>
                    <FormattedMessage id="common.home.features.portfolio.title" />
                </h3>
                <p>
                    <FormattedMessage id="common.home.features.portfolio.paragraph" />
                </p>

                <Link to={getEndpointPath(commonPortfolioHomeEndpoint, null)}
                    className="feature__btn">
                    <FormattedMessage id="common.home.features.portfolio.button" />
                </Link>
            </div>

            <div className="page-home__hero-features__feature">
                <h3>
                    <FormattedMessage id="common.home.features.cheatsheet.title" />
                </h3>
                <p>
                    <FormattedMessage id="common.home.features.cheatsheet.paragraph" />
                </p>

                <Link to={getEndpointPath(cheatsheetHomeEndpoint, null)}
                    className="feature__btn">
                    <FormattedMessage id="common.home.features.cheatsheet.button" />
                </Link>
            </div>
        </main>
        <section className="page-home__hero-contact">
            <div>
                <h3>
                    <FormattedMessage id="common.home.contact.title" />
                </h3>
                <p>
                    <FormattedMessage id="common.home.contact.paragraph" />
                </p>
                <div className="hero-contact__container">
                    <div className="hero-contact__container__phone">
                        <div>
                            <a href={`tel:${intl.formatMessage({ id: "common.contact.phone.raw" })}`}>
                                <ImageUtil
                                    src={phoneImage}
                                />
                            </a>
                        </div>
                        <div>
                            <a href={`tel:${intl.formatMessage({ id: "common.contact.phone.raw" })}`}>
                                <FormattedMessage id="common.contact.phone" />
                            </a>
                        </div>

                    </div>
                    <div className="hero-contact__container__email">
                        <div>
                            <a href={`mailto:${intl.formatMessage({ id: "common.contact.email.raw" })}`}>
                                <ImageUtil
                                    src={emailImage}
                                />
                            </a>
                        </div>
                        <div>
                            <a href={`mailto:${intl.formatMessage({ id: "common.contact.email.raw" })}`}>
                                <FormattedMessage id="common.contact.email" />
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
}