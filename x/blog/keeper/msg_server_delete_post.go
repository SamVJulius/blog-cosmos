package keeper

import (
    "context"
    "fmt"

    "blog/x/blog/types"

    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    val, found := k.GetPost(ctx, msg.Id)
    if !found {
        return nil, fmt.Errorf("key %d doesn't exist: %w", msg.Id, sdkerrors.ErrKeyNotFound)
    }
    if msg.Creator != val.Creator {
        return nil, fmt.Errorf("incorrect owner: %w", sdkerrors.ErrUnauthorized)
    }
    k.RemovePost(ctx, msg.Id)
    return &types.MsgDeletePostResponse{}, nil
}