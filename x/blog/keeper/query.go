package keeper

import (
	"context"

	"blog/x/blog/types"
)

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// ListPost implements [types.QueryServer].
func (q queryServer) ListPost(ctx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {
	return q.k.ListPost(ctx, req)
}

// ShowPost implements [types.QueryServer].
func (q queryServer) ShowPost(ctx context.Context, req *types.QueryShowPostRequest) (*types.QueryShowPostResponse, error) {
	return q.k.ShowPost(ctx, req)
}