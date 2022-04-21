import { Game, makeFunctonalGame } from "@app/util/game/Game"
import { shuffle } from "@app/util/lang/shuffle"
import { BrowseWordsGameConfig, BrowseWordsGameInput, BrowseWordsGameState } from "./Define"

export const browseWordsGameFactory = (config: BrowseWordsGameConfig): Game<BrowseWordsGameState, BrowseWordsGameInput> => {
    const tuples = [...config.tuples]
    if (config.sortMode === "shuffle") {
        shuffle(tuples)
    } else if (config.sortMode === "alphabetic") {
        tuples.sort((a, b) => a.sourceWord.localeCompare(b.sourceWord))
    } else {
        throw new Error(`Unknown sort mode ${config.sortMode}`)
    }

    const initialState: BrowseWordsGameState = {
        isFinalized: config.tuples.length === 0,

        tupleIndex: 0,
        tupleIteration: 0,
        iteration: 0,

        tuples,
        showTuple: false,

        sortMode: config.sortMode,
    }

    return makeFunctonalGame(initialState, async (state, input) => {
        if (input.type === "next-word") {
            let tuples = state.tuples
            let newIndex = state.tupleIndex + 1
            let tupleIteration = state.tupleIteration
            if (newIndex >= state.tuples.length) {
                newIndex = 0
                if (state.sortMode === "shuffle") {
                    tuples = [...tuples]
                    shuffle(tuples)
                }

                tupleIteration++
            }
            return {
                ...state,
                tuples,
                tupleIndex: newIndex,
                showTuple: false,
                tupleIteration,
                iteration: state.iteration + 1,
            }
        } else if (input.type === "show") {
            return {
                ...state,
                showTuple: true,
            }
        } else {
            throw new Error("Unreachable code")
        }
    })
}