package cosmicproof

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "github.com/atahanyild/ytubc-lambda/x/cosmicproof/keeper"
    "github.com/atahanyild/ytubc-lambda/x/cosmicproof/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
    return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
        switch msg := msg.(type) {
        case *types.MsgCreatecosmicproof:
            return handleMsgCreatecosmicproof(ctx, k, msg)
        default:
            return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown message type")
        }
    }
}

func handleMsgCreatecosmicproof(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatecosmicproof) (*sdk.Result, error) {
    err := k.Createcosmicproof(ctx, msg)
    if err != nil {
        return nil, err
    }
    return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
