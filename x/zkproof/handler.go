package zkproof

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "github.com/atahanyild/ytubc-lambda/x/zkproof/keeper"
    "github.com/atahanyild/ytubc-lambda/x/zkproof/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
    return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
        switch msg := msg.(type) {
        case *types.MsgCreateZKProof:
            return handleMsgCreateZKProof(ctx, k, msg)
        default:
            return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown message type")
        }
    }
}

func handleMsgCreateZKProof(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateZKProof) (*sdk.Result, error) {
    err := k.CreateZKProof(ctx, msg)
    if err != nil {
        return nil, err
    }
    return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
