package keeper

import (
	"errors"
	"github.com/VelaChain/orion/x/chios/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatePool(pool types.Pool) bool {
	return pool.Validate()
}

func (k Keeper) SetPool(ctx sdk.Context, pool *types.Pool) error {
	if !pool.Validate(){
		// TODO add to errors
		return errors.New("pool invalid")
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolKeyFromPoolName(pool.PoolId)
	store.Set(key, k.cdc.MustMarshal(pool))
	return nil
}

func (k Keeper) GetPool(ctx sdk.Context, poolName string) (types.Pool, error) {
	var pool types.Pool
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolKeyFromPoolName(poolName)
	if !k.Exists(ctx, key) {
		// TODO add to errors
		return pool, errors.New("Pool DNE")
	}
	bz := store.Get(key)
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, nil
}

func (k Keeper) RemovePool(ctx sdk.Context, poolName string) errpr {
	store := ctx.KVStore(k.storeKey)
	key := type.GetPoolKeyFromPoolName(poolName)
	if !k.Exists(ctx, key) {
		// TODO add to errors
		return errors.New("Pool DNE")
	}
		
}

func (k Keeper) GetPools(ctx sdk.Context) []*types.Pool {
	var pools []*types.Pool
	Iterator := k.GetPoolsInterator(ctx)
	defer func(iterator sdk.Iterator){
		err := iterator.Close()
		if err != nil {
			panic(err)
		}
	}(iterator)
	for ; iterator.Valid(); iterator.Next() {
		var pool types.Pool
		bytesVal := iterator.Value()
		k.cdc.MustUnmarshal(bytesVal, &pool)
		pools = append(pools, &pool)
	}
	return pools
}

func (k Keeper) GetPoolsPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]*types.Pool, *query.PageResponse, error) {
	var pools []*type.Pool
	store := ctx.KVStore(k.storeKey)
	poolStore := prefix.NewStore(store, type.KeyPoolPrefix)
	pageRes, err := query.Paginate(poolStore, pagination, func (key []byte, value []byte) error {
		var pool types.Pool
		err := k.cdc.Unmarshal(value, &pool)
		if err != nil {
			return err
		}
		pools = append(pools, &pool)
		return nil
	})
	if err != nil {
		return nil, &query.PageResponse{}, status.Error(codes.Internal, err.Error())
	}
	return pools, pageRes, nil
}

func (k Keeper) GetPoolsInterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.KeyPoolPrefix)
}