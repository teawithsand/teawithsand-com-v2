import { BaseGameState } from "@app/util/game/Game";
import { WordTuple } from "../../api/entity/WordTuple";

export interface BrowseWordsGameConfig {
    tuples: WordTuple[],
    sortMode: "shuffle" | "alphabetic",
}

export interface BrowseWordsGameState extends BaseGameState {
    tuples: WordTuple[],
    tupleIndex: number,
    showTuple: boolean,

    tupleIteration: number,
    iteration: number,

    sortMode: "shuffle" | "alphabetic",
}

export type BrowseWordsGameInput = {
    type: "next-word",
} | {
    type: "show",
}
