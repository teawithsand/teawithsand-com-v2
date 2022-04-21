package langka

import "github.com/teawithsand/webpage/util/voter"

const VoteCreateWordSet = voter.Operation("VOTE_LANGKA_WORD_SET_CREATE")
const VoteEditWordSet = voter.Operation("VOTE_LANGKA_WORD_SET_EDIT")
const VoteDeleteWordSet = voter.Operation("VOTE_LANGKA_WORD_SET_DELETE")
const VotePublishWordSet = voter.Operation("VOTE_LANGKA_WORD_SET_PUBLISH")
const VoteUnpublishWordSet = voter.Operation("VOTE_LANGKA_WORD_SET_UNPUBLISH")
const VoteGetSecretWordSet = voter.Operation("VOTE_LANGKA_WORD_SET_GET_SECRET")
const VoteGetPublicWordSet = voter.Operation("VOTE_LANGKA_WORD_SET_GET_PUBLIC")
const VoteGetOwnedWordSetList = voter.Operation("VOTE_LANGKA_WORD_SET_GET_OWNED_LIST")
