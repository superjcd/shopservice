package store

import (
	"context"

	v1 "github.com/superjcd/shopservice/genproto/v1"
)

type Factory interface {
	Shops() ShopStore
	Close() error
}

type ShopStore interface {
	Create(ctx context.Context, _ *v1.CreateShopRequest) error
	List(ctx context.Context, _ *v1.ListShopRequest) (*ShopList, error)
	Update(ctx context.Context, _ *v1.UpdateShopRequest) error
	Delete(ctx context.Context, _ *v1.DeleteShopRequest) error
}
