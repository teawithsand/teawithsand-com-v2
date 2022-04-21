import ImageUtil from "@app/util/react/image/ImageUtil"
import React from "react"
import { FormattedMessage } from "react-intl"

import loading from "@app/images/loading.svg"

export default () => {
    return <div className="app-spinner">
        <ImageUtil src={loading} className="app-spinner__image" />
        <div className="app-spinner__text">
            <FormattedMessage id="common.loading_spinner.loading" />
        </div>
    </div>
}