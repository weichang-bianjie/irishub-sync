package nft

import (
	. "github.com/irisnet/irishub-sync/util/constant"
	"strings"
	"github.com/irisnet/irishub-sync/store/document"
	"github.com/irisnet/irishub-sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irishub-sync/store"
)

type (
	DocMsgNFTTransfer struct {
		Sender    string `bson:"sender"`
		Recipient string `bson:"recipient"`
		URI       string `bson:"uri"`
		Denom     string `bson:"denom"`
		Id        string `bson:"id"`
		Data      string `bson:"data"`
	}
)

func (m *DocMsgNFTTransfer) Type() string {
	return TxTypeNFTTransfer
}

func (m *DocMsgNFTTransfer) BuildMsg(v interface{}) {
	msg := v.(*types.MsgTransferNFT)

	m.Sender = msg.Sender.String()
	m.Recipient = msg.Recipient.String()
	m.Id = strings.ToLower(msg.Id)
	m.Denom = strings.ToLower(msg.Denom)
	m.URI = msg.URI
	m.Data = msg.Data
}

func (m *DocMsgNFTTransfer) HandleTxMsg(msgData sdk.Msg, tx *document.CommonTx) *document.CommonTx {

	m.BuildMsg(msgData)
	tx.Msgs = append(tx.Msgs, document.DocTxMsg{
		Type: m.Type(),
		Msg:  m,
	})
	tx.Addrs = append(tx.Addrs, m.Sender, m.Recipient)
	tx.Types = append(tx.Types, m.Type())
	if len(tx.Msgs) > 1 {
		return tx
	}
	tx.Type = m.Type()
	tx.From = m.Sender
	tx.To = m.Recipient
	tx.Amount = []store.Coin{}
	return tx
}
