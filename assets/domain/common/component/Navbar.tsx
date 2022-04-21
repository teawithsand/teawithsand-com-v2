import { Nav, Navbar, NavDropdown } from "react-bootstrap"
import React, { useMemo } from "react"
import logo from "@app/images/tws.png"
import { LinkContainer } from "react-router-bootstrap"
import { FormattedMessage, useIntl } from "react-intl"
import ImageUtil from "@app/util/react/image/ImageUtil"
import { Link } from "react-router-dom"
import { getEndpointPath, commonHomeEndpoint, userRegisterEndpoint, langkaHomeEndpoint, userLoginEndpoint, useEndpointNavigate, userSecretProfileEndpoint } from "../endpoints"
import { useDispatch, useSelector } from "react-redux"
import { unsetUnsetDataAction, userDataSelector } from "../redux/user"

interface NavbarEntryBase {
    messageId: string,
    values?: { [key: string]: string },
}

interface NavbarDataEntry extends NavbarEntryBase {
    path: string,
    entries?: undefined,
    onClick?: undefined,
}

interface ClickableDataEntry extends NavbarEntryBase {
    onClick: () => void,
    entries?: undefined,
    path?: undefined,
}

interface NavbarCollapseEntry extends NavbarEntryBase {
    entries: InnerNavbarEntry[],
    onClick?: undefined,
    path?: undefined,
}

type InnerNavbarEntry = NavbarDataEntry | ClickableDataEntry
type NavbarEntry = InnerNavbarEntry | NavbarCollapseEntry


const InnerEntry = (props: { entry: InnerNavbarEntry, dropdown?: boolean }) => {
    const intl = useIntl()

    const { entry, dropdown } = props

    if (typeof entry.path !== "undefined") {
        if (dropdown) {
            return <LinkContainer to={entry.path}>
                <NavDropdown.Item href="#">
                    <FormattedMessage id={entry.messageId} values={entry.values ?? {}} />
                </NavDropdown.Item>
            </LinkContainer>
        } else {
            return <Nav>
                <LinkContainer to={entry.path}>
                    <Nav.Link>
                        <FormattedMessage id={entry.messageId} values={entry.values ?? {}} />
                    </Nav.Link>
                </LinkContainer>
            </Nav>
        }

    } else if (typeof entry.onClick !== "undefined") {
        if (dropdown) {
            return <NavDropdown.Item onClick={() => {
                entry.onClick()
            }}>
                <FormattedMessage id={entry.messageId} values={entry.values ?? {}} />
            </NavDropdown.Item>
        } else {
            return <Nav>
                <Nav.Link onClick={() => {
                    entry.onClick()
                }}>
                    <FormattedMessage id={entry.messageId} values={entry.values ?? {}} />
                </Nav.Link>
            </Nav>
        }
    }
}

const Entry = (props: { entry: NavbarEntry }) => {
    const intl = useIntl()

    const { entry } = props

    if (typeof entry.entries !== "undefined") {
        return <Nav>
            <NavDropdown title={intl.formatMessage({ id: entry.messageId }, entry.values ?? {})}>
                {entry.entries.map((ie, j) => {
                    return <InnerEntry entry={ie} key={j} dropdown />
                })}
            </NavDropdown>
        </Nav>
    } else if (typeof entry.path !== "undefined") {
        return <Nav>
            <LinkContainer to={entry.path}>
                <Nav.Link>
                    <FormattedMessage id={entry.messageId} values={entry.values ?? {}} />
                </Nav.Link>
            </LinkContainer>
        </Nav>
    } else if (typeof entry.onClick !== "undefined") {
        return <Nav>
            <Nav.Link onClick={() => {
                entry.onClick()
            }}>
                <FormattedMessage id={entry.messageId} values={entry.values ?? {}} />
            </Nav.Link>
        </Nav>
    }
}

const Entries = (props: { entries: NavbarEntry[] }) => {
    const { entries: navbarEntries } = props

    return <>
        {
            navbarEntries.map((e, i) => <Entry entry={e} key={i} />)
        }
    </>
}

// TODO(teawithsand): rewrite with bootstrap-less style
export default () => {
    const intl = useIntl()

    const userData = useSelector(userDataSelector)
    const dispatch = useDispatch()
    const navigate = useEndpointNavigate()

    const leftEntries: NavbarEntry[] = useMemo(() => [
        {
            messageId: "common.navbar.about_me",
            path: "/about-me",
        },
        {
            messageId: "common.navbar.langka_home",
            path: getEndpointPath(langkaHomeEndpoint, null),
        }
    ], [])

    const rightEntries: NavbarEntry[] = useMemo(() => [
        ...(userData ? [
            {
                messageId: "common.navbar.logged_in",
                values: {
                    "publicName": userData.publicName,
                },

                entries: [
                    {
                        messageId: "common.navbar.logout",
                        onClick: () => {
                            dispatch(unsetUnsetDataAction())
                            navigate(commonHomeEndpoint, null)
                        },
                    },
                    {
                        messageId: "common.navbar.my_profile",
                        path: getEndpointPath(userSecretProfileEndpoint, {
                            id: userData.id,
                        }),
                    },
                ]
            },
        ] : [
            {
                messageId: "common.navbar.register",
                path: getEndpointPath(userRegisterEndpoint, null),
            },
            {
                messageId: "common.navbar.login",
                path: getEndpointPath(userLoginEndpoint, null),
            },
        ])
    ], [userData])


    return <>
        <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark" sticky="top" className="ps-1 pe-1 ps-md-5 pe-md-5">
            <Navbar.Brand>
                <Link to="/">
                    <ImageUtil
                        src={logo}
                        className="d-inline-block align-top navbar--logo"
                        alt={
                            intl.formatMessage({ id: "common.navbar.logo.alt" })
                        } />
                </Link>
            </Navbar.Brand>

            <Nav className="d-inline-block ms-auto me-auto">
                <LinkContainer to={getEndpointPath(commonHomeEndpoint, null)}>
                    <Nav.Link>
                        <FormattedMessage id="common.navbar.home" />
                    </Nav.Link>
                </LinkContainer>
            </Nav>

            <Navbar.Toggle />

            <Navbar.Collapse>
                <Entries entries={leftEntries} />
                <div className="ms-auto"></div>
                <Entries entries={rightEntries} />
            </Navbar.Collapse>


        </Navbar>
    </>
}