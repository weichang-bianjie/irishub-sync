package evidence

import (
	"github.com/irisnet/irishub-sync/store/document"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irishub-sync/types"
)

func HandleTxMsg(msg sdk.Msg, tx *document.CommonTx) (*document.CommonTx, bool) {
	ok := true
	switch msg.Type() {
	case new(types.MsgSubmitEvidence).Type():
		docMsg := DocMsgSubmitEvidence{}
		return docMsg.HandleTxMsg(msg, tx), ok
	default:
		ok = false
	}
	return tx, ok
}