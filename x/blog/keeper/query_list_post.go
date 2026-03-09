package keeper

import (
	"context"

	"blog/x/blog/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPost(goCtx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var posts []types.Post
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)), types.KeyPrefix(types.PostKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &post)
		posts = append(posts, post)
	}

	return &types.QueryListPostResponse{
		Post: posts,
	}, nil
}
