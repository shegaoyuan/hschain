// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: hschain/x/bank/types
package bank

import (
	"github.com/hschain/hschain/x/bank/internal/keeper"
	"github.com/hschain/hschain/x/bank/internal/types"
)

const (
	DefaultCodespace         = types.DefaultCodespace
	CodeSendDisabled         = types.CodeSendDisabled
	CodeInvalidInputsOutputs = types.CodeInvalidInputsOutputs
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	QuerierRoute             = types.QuerierRoute
	DefaultParamspace        = types.DefaultParamspace
	DefaultSendEnabled       = types.DefaultSendEnabled

	EventTypeTransfer      = types.EventTypeTransfer
	AttributeKeyRecipient  = types.AttributeKeyRecipient
	AttributeKeySender     = types.AttributeKeySender
	AttributeValueCategory = types.AttributeValueCategory
)

var (
	// functions aliases
	RegisterCodec          = types.RegisterCodec
	ErrNoInputs            = types.ErrNoInputs
	ErrNoOutputs           = types.ErrNoOutputs
	ErrInputOutputMismatch = types.ErrInputOutputMismatch
	ErrSendDisabled        = types.ErrSendDisabled
	NewBaseKeeper          = keeper.NewBaseKeeper
	NewInput               = types.NewInput
	NewOutput              = types.NewOutput
	ParamKeyTable          = types.ParamKeyTable

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	ParamStoreKeySendEnabled = types.ParamStoreKeySendEnabled
)

type (
	BaseKeeper   = keeper.BaseKeeper // ibc module depends on this
	Keeper       = keeper.Keeper
	MsgSend      = types.MsgSend
	MsgMultiSend = types.MsgMultiSend
	Input        = types.Input
	Output       = types.Output
)
