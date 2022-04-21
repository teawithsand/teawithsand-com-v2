import { GameRenderer, GameRendererProps } from "@app/util/game/GameRenderer";
import { Game, GameHelperCallbacks, useGame } from "@app/util/game/Game";

export default <S, I, C>(props: {
    gameFactory: () => Game<S, I>,
    renderer: GameRenderer<S, I>,

    callbacks?: GameHelperCallbacks<S, I>,
}) => {
    const { gameFactory, renderer, callbacks } = props
    const [state, onInput] = useGame({
        callbacks: callbacks ?? {},
        gameFactory,
    }, [renderer])

    const rendererProps: GameRendererProps<S, I> = {
        handleInput: onInput,

        state: state.state,
        isFinalized: state.type === "finalized",

        hasError: state.type === "error",
        error: state.type === "error" ? state.error : undefined,

        isComputingNewState: state.type === "computingNewState",
    }

    return renderer(rendererProps)
}