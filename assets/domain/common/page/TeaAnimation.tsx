import classnames from '@app/util/lang/classnames';
import { useBreakpointIndex } from '@app/util/react/hook/breakpointHook';
import useCss from '@app/util/react/hook/useCss';
import useCssVar from '@app/util/react/hook/useCssVar';
import useWindowDimensions from '@app/util/react/hook/windowDimesionsHook';
import React, { useMemo } from 'react';

require("@app/newstyles/common/component/tea-animation.scss")

// TODO(teawithsand): here create animated flowing tea using webgl and shaders
// since anime.js is too slow
export default (props: {
    className?: string,
}) => {
    const range = (n: number) => Array.from(Array(n).keys())

    const maxColumn = parseInt(useCssVar("--app-tea-columns", []))
    const maxSand = parseInt(useCssVar("--app-tea-styles", []))

    const br = useBreakpointIndex()
    let auxSandCount = useMemo(() => {
        if (br <= 1) {
            return 100
        } else if (br >= 4) {
            return 400
        } else {
            return 250
        }
    }, [br])

    auxSandCount = Math.max(0, auxSandCount - maxColumn * 2)

    const randomSand = () => Math.round(Math.random() * 1000000) % maxSand + 1
    const randomColumn = () => Math.round(Math.random() * 1000000) % maxColumn + 1

    const randomSands = useMemo(() => {
        return [
            range(maxColumn).map(() => randomSand()),
            range(maxColumn).map(() => randomSand()),
            range(auxSandCount).map(() => randomSand()),
        ]
    }, [auxSandCount])

    const randomColumns = useMemo(() => {
        return range(auxSandCount).map(() => randomColumn())
    }, [auxSandCount])

    return <div className={classnames("app-tea", props.className)}>
        {
            range(maxColumn)
                .map((i) => <div
                    key={i}
                    className={`app-tea__sand app-tea__sand--flow-${randomSands[0][i]} app-tea__sand--col-${i + 1}`}
                ></div>)
        }
        {
            range(maxColumn)
                .map((i) => <div
                    key={i}
                    className={`app-tea__sand app-tea__sand--flow-${randomSands[1][i]} app-tea__sand--col-${i + 1}`}
                ></div>)
        }
        {
            range(auxSandCount)
                .map((i) => <div
                    key={i}
                    className={`app-tea__sand app-tea__sand--flow-${randomSands[2][i]} app-tea__sand--col-${randomColumns[i]}`}
                ></div>)
        }
    </div >
}