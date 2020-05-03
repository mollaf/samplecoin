package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mollaf/samplecoin/x/nameservice/types"
)

// NewHandler returns a hhandler for "nameservice" type messages
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nameservice Msg type %v", msg.Type()))
		}
	}
}

// Handle message to set name
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) (*sdk.Result, error) {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	keeper.SetName(ctx, msg.Name, msg.Value)
	return &sdk.Result{}, nil
}

// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) (*sdk.Result, error) {

	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
	}
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
		if err != nil {
			return nil, err
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return &sdk.Result{}, nil
}

// Handle a message to delete name
func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg MsgDeleteName) (*sdk.Result, error) {
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.Name)
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.DeleteWhois(ctx, msg.Name)
	return &sdk.Result{}, nil
}

// // NewHandler creates an sdk.Handler for all the nameservice type messages
// func NewHandler(k Keeper) sdk.Handler {
// 	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
// 		ctx = ctx.WithEventManager(sdk.NewEventManager())
// 		switch msg := msg.(type) {
// 		// TODO: Define your msg cases
// 		//
// 		//Example:
// 		// case Msg<Action>:
// 		// 	return handleMsg<Action>(ctx, k, msg)
// 		default:
// 			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName,  msg)
// 			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
// 		}
// 	}
// }
//
// // handle<Action> does x
// func handleMsg<Action>(ctx sdk.Context, k Keeper, msg Msg<Action>) (*sdk.Result, error) {
// 	err := k.<Action>(ctx, msg.ValidatorAddr)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// TODO: Define your msg events
// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
// 			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
// 		),
// 	)
//
// 	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
// }
