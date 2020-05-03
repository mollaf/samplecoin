package nameservice

import (
	"github.com/mollaf/samplecoin/x/nameservice/keeper"
	"github.com/mollaf/samplecoin/x/nameservice/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases

	NewMsgSetName    = types.NewMsgSetName
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgSetName    = types.MsgSetName
	MsgBuyName    = types.MsgBuyName
	MsgDeleteName = types.MsgDeleteName
	Whois         = types.Whois

	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	// TODO: Fill out module types
)
