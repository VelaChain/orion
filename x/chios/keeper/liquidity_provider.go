package keeper

import (
	"github.com/VelaChain/orion/x/chios/type"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetLiqProv(ctx sdk.Context, lp *types.LiquidityProvider) {
	if !lp.Validate() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetProviderKey(lp.Liquidity.Symbol, lp.Creator)
	store.Set(key, k.cdc.MustMarshal(lp))
}

func (k Keeper) GetLiqProv(ctx sdk.Context, poolName string, lpAddr string) (types.LiquidityProvider, error) {
	var lp types.LiquidityProvider
	key := types.GetProviderKey(poolName, lpAddr)
	store := ctx.KVStore(k.storeKey)
	if !k.Exists(ctx, key) {
		// TODO add to errors
		return lp, errors.New("LP DNE")
	}

	bz := store.Get(key)
	k.cdc.MustUnmarshal(bz, &lp)
	return lp, nil
}

func (k Keeper) RemoveLiqPro(ctx sdk.Context, poolName string, lpAddr string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProviderKey(poolName, lpAddr)
	if !k.Exists(ctx, key) {
		// TODO add to errors
		return lp, errors.New("LP DNE")
	}

	bz := store.Get(key)
	k.cdc.MustUnmarshal(bz, &lp)
	if lp.Liquidity.IsNegative() {
		// TODO add to errors
		return lp, errors.New("Provider has negative liquidity")
	}
	if !lp.Liquidity.IsZero() {
		// TODO add to errors
		return lp, errors.New("Provider still has liquidity")
	}
	
	store.Delete(key)
}

// returns providers for a given pool
func (k Keeper) GetProviders(ctx sdk.Context, poolName string) []*types.LiquidityProvider, error {
	var providers []*types.LiquidityProvider
	iterator := k.GetProviderIterator(ctx, poolName)
	defer func(iterator sdk.Iterator){
		err := iterator.Close()
		if err != nil {
			panic(err)
		}
	}(iterator)
	for ; iterator.Valid(); iterator.Next(){
		var lp types.LiquidityProvider
		bytesVal := iterator.Value()
		k.cdc.MustUnmarshal(bytesVal, &lp)
		if lp.Liquidity.Symbol != poolName {
			// TODO add to errors
			return providers, errors.New("Iterator returned provider from different pool")
		}
		providers = append(providers, &lp)
	}
	return providers, nil
}

func (k Keeper) GetAllProviders(ctx sdk.Context) []*types.LiquidityProvider {
	var providers []*types.LiquidityProvider
	iterator := k.GetAllProvidersIterator(ctx)
	defer func(iterator sdk.Iterator){
		err := iterator.Close()
		if err != nil {
			panic(err)
		}
	}(iterator)
		var lp types.LiquidityProvider
		bytesVal := iterator.Value()
		k.cdc.MustUnmarshal(bytesVal, &lp)
		providers = append(providers, &lp)
	}
	return providers
}

// returns providers from a given pool
func (k Keeper) GetProvidersPaginated(ctx sdk.Context, poolName string, pagination *query.PageRequest) ([]*types.LiquidityProvider, *query.PageResponse, error) {
	var providers []*types.LiquidityProvider
	store := ctx.KVStore(k.storeKey)
	provStore := prefix.NewStore(store, types.GetProvidersKey(poolName))
	pageRes, err := query.Paginate(provStore, pagination, func(key []byte, value []byte){
		var  lp types.LiquidityProvider
		err := k.cdc.Unmarshal(value, &lp)
		if err != nil {
			return err
		} 
		providers = append(providers, &lp)
		return nil
	})
	if err != nil {
		return nil, &query.PageResponse{}, status.Error(codes.Internal, err.Error())
	}
	return providers, pageRes, nil
}

func (k Keeper) GetAllProvidersPaginated(ctx sdk.Context, pagination *query.PageRequest) ([]*types.LiquidityProvider, *query.PageResponse, error) {
	var providers []*types.LiquidityProvider
	store := ctx.KVStore(k.storeKey)
	provStore := prefix.NewStore(store, types.KeyProviderPrefix)
	pageRes, err := query.Paginate(provStore, pagination, func(key []byte, value []byte){
		var lp types.LiquidityProvider
		err := k.cdc.Unmarshal(value, &lp)
		if err != nil {
			return err
		}
		providers = append(providers, &lp)
		return nil
	})
	if err != nil {
		return nil, &query.PageResponse{}, status.Error(codes.Internal, err.Error())
	}
	return providers, pageRes, nil
}

// iterator for all providers for all pools
func (k Keeper) GetAllProvidersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.KeyProviderPrefix)
}

// iterator for all providers for a specific pool
func (k Keeper) GetProvidersIterator(ctx sdk.Context, poolName string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.GetProvidersKey(poolName))
}

