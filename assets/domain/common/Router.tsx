import { Routes, Route, HashRouter } from 'react-router-dom';
import React from "react"
import CommonHome from './page/CommonHome';
import { commonAboutMeEndpoint, getEndpointRoute, commonHomeEndpoint, langkaHomeEndpoint, userRegisterEndpoint, Endpoint, userLoginEndpoint, userSecretProfileEndpoint, userChangePasswordEndpoint, userChangeEmailEndpoint, cheatsheetHomeEndpoint, langkaCreateWordSetEndpoint, langkaEditWordSetEndpoint, langkaListPublicWordSetsEndpoint, langkaListOwnedWordSetsEndpoint } from './endpoints';
import Navbar from './component/Navbar';
import Footer from './component/Footer';
import LangkaHome from '../langka/page/LangkaHome';
import UserRegister from '../user/page/UserRegister';
import UserLogin from '../user/page/UserLogin';
import UserSecretProfile from '../user/page/UserSecretProfile';
import UserChangePassword from '../user/page/UserChangePassword';
import UserChangeEmail from '../user/page/UserChangeEmail';
import CheatsheetHome from '../cheatsheet/page/CheatsheetHome';
import WordSetCreatePage from '../langka/page/WordSetCreatePage';
import AboutMePage from './page/about/AboutMePage';
import WordSetEditPage from '../langka/page/WordSetEditPage';
import WordSetListPublic from '../langka/page/WordSetListPublic';
import WordSetListOwned from '../langka/page/WordSetListOwned';


interface Route {
    endpoint: Endpoint<any>,
    component: any,
}

const endpoints: Route[] = [
    { endpoint: commonHomeEndpoint, component: CommonHome },
    { endpoint: commonAboutMeEndpoint, component: AboutMePage, },

    { endpoint: userRegisterEndpoint, component: UserRegister },
    { endpoint: userLoginEndpoint, component: UserLogin },
    { endpoint: userSecretProfileEndpoint, component: UserSecretProfile },
    { endpoint: userChangePasswordEndpoint, component: UserChangePassword },
    { endpoint: userChangeEmailEndpoint, component: UserChangeEmail },

    { endpoint: langkaHomeEndpoint, component: LangkaHome },
    { endpoint: langkaCreateWordSetEndpoint, component: WordSetCreatePage },
    { endpoint: langkaEditWordSetEndpoint, component: WordSetEditPage },
    { endpoint: langkaListPublicWordSetsEndpoint, component: WordSetListPublic },
    { endpoint: langkaListOwnedWordSetsEndpoint, component: WordSetListOwned },

    { endpoint: cheatsheetHomeEndpoint, component: CheatsheetHome },
]


export const AppRouter = () => {
    return <>
        <HashRouter>
            <Navbar />
            <Routes>
                {
                    endpoints
                        .map(r => ({
                            ...r,
                            path: getEndpointRoute(r.endpoint),
                        }))
                        .map(({ path, component: Component }) =>
                            <Route key={path} path={path} element={<Component />} />
                        )
                }

                {/* TODO(teawithsand): not found route */}
            </Routes>
            <Footer />
        </HashRouter>
    </>
}