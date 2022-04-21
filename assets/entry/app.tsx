import { render } from "react-dom"
import React from "react"
import { IntlProvider } from "react-intl"
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';
import { i18nConfig } from "@app/util/i18n/i18n"
import { Store } from "redux"
import { Persistor } from "redux-persist"

import "@app/styles/app.scss"

import { applyIntlErrorHack } from "@app/util/i18n/init";
import { AppRouter } from "@app/domain/common/Router";
import { store, persistor } from "@app/domain/common/redux/store";

const App = (props: {
    store: Store,
    persistor: Persistor,
}) => {
    return <PersistGate loading={null} persistor={props.persistor}>
        <React.StrictMode>
            <IntlProvider
                locale={i18nConfig.locale}
                defaultLocale={i18nConfig.locale}
                messages={i18nConfig.messages}
            >
                <Provider store={props.store}>
                    <AppRouter />
                </Provider>
            </IntlProvider>
        </React.StrictMode>
    </PersistGate>
}

const initApp = async () => {
    applyIntlErrorHack()

    const target = document.createElement("div")
    target.id = "root"
    document.body.appendChild(target)

    render(<App persistor={persistor} store={store} />, document.getElementById('root'));
}

document.title = "teawithsand.com"

document.addEventListener("DOMContentLoaded", () => {
    initApp()
})


// disable SW for dev, since it's caching is annoying and sometimes causes old versions of code to be loaded
// it integrates badly with hot reload
window.addEventListener("load", () => {
    // disable SW for the time being 
    // Website is fast enough anyway
    // loadAndRegisterServiceWorker()
})