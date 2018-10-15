package common

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAccountCreate] = func() types.Operation {
		op := &SignMessageOperation{}
		return op
	}
}

type SignMessageOperation struct {
	types.OperationFee
	From types.GrapheneID `json:"from"`
	To types.GrapheneID `json:"to"`
	Proxy_account types.GrapheneID `json:"proxy_account"`
	Amount types.Asset `json:"amount"`
	Percent types.Int16 `json:"percent"`
	Memo string `json:"memo"`
	Expiration types.Time `json:"expiration"`
	Extension uint8 `json:"extension"`
}

func (p SignMessageOperation) Type() types.OperationType{
	return types.OperationTypeAccountCreate
}


func (p SignMessageOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode from")
	}

	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}

	if err := enc.Encode(p.Proxy_account); err != nil {
		return errors.Annotate(err, "encode proxy account")
	}

	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}

	if err := enc.Encode(p.Percent); err != nil {
		return errors.Annotate(err, "encode percent")
	}

	if err := enc.Encode(p.Memo); err != nil {
		return errors.Annotate(err, "encode memo")
	}

	if err := enc.Encode(p.Expiration); err != nil {
		return errors.Annotate(err, "encode expiration")
	}

	if err := enc.Encode(p.Extension); err != nil {
		return errors.Annotate(err, "encode extension")
	}

	return nil
}

//NewAccountCreateOperation creates a new AccountCreateOperation
func NewSignMessageOperation() *SignMessageOperation {
	tx := SignMessageOperation{
		Extension: 8,
	}
	return &tx
}