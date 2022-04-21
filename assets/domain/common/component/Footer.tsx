import { Container, } from "react-bootstrap"
import * as React from "react"
import { FormattedMessage } from "react-intl"

export default () => {
    return <footer className="mt-auto mb-5">
        <hr></hr>
        <Container className="text-end text-muted">
            <FormattedMessage id="common.footer.text" />
        </Container>
    </footer>
}