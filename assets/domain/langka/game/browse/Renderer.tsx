import { GameRenderer } from "@app/util/game/GameRenderer";
import { BrowseWordsGameInput, BrowseWordsGameState } from "./Define";
import React, { useMemo } from "react"
import { FormattedMessage } from "react-intl";

export interface BrowseWordsGameRendererOptions {

}

export const browseWordsGameRendererFactory =
    (options: BrowseWordsGameRendererOptions): GameRenderer<BrowseWordsGameState, BrowseWordsGameInput> => ({ state, handleInput }) => {
        const tuple = state.tuples[state.tupleIndex]

        const ShowButton = () => <button className="btn btn-primary" onClick={() => {
            handleInput({
                type: "show"
            })
        }}>
            <FormattedMessage id="langka.browsewordsgame.show" />
        </button>

        const NextButton = () => <button className="btn btn-primary" onClick={() => {
            handleInput({
                type: "next-word"
            })
        }}>
            <FormattedMessage id="langka.browsewordsgame.next-word" />
        </button>

        return <div
            className="browsewords-game">
            <div className="browsewords-game__game-summary">
                <h1>
                    <FormattedMessage id="langka.browsewordsgame.title" defaultMessage={"Browse words"} />
                </h1>
                <p>
                    <FormattedMessage id="langka.browsewordsgame.state.iteration" values={{
                        iteration: Math.floor(state.iteration / state.tuples.length) + 1,
                    }} defaultMessage={"Iteration #{iteration}"} />
                </p>
                <p>
                    <FormattedMessage id="langka.browsewordsgame.state.position" values={{
                        iteration: state.tupleIndex + 1,
                        count: state.tuples.length,
                    }} defaultMessage={"Word #{iteration} out of {count}"} />
                </p>
            </div>
            <div className="browsewords-game__source-word">
                {tuple.sourceWord}
            </div>
            <div className="browsewords-game__destination-words">
                {state.showTuple ? tuple.destinationWords.join(" ") : <FormattedMessage id="langka.browsewordsgame.destination.hidden" />}
            </div>
            <div className="browsewords-game__description">
                <p>
                    {state.showTuple ? tuple.description : <FormattedMessage id="langka.browsewordsgame.description.hidden" />}
                </p>
            </div>
            <div className="browsewords-game__controls">
                {state.showTuple ? <NextButton /> : <ShowButton />}
            </div>
        </div>
    }