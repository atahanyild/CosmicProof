package zkproof

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/your/module/x/zkproof/keeper"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
    return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
        switch msg := msg.(type) {
        // Handle your module messages here
        default:
            errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
            return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
        }
    }
}
