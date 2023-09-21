//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/superjcd/shopservice/config"
	v1 "github.com/superjcd/shopservice/genproto/v1"
	"github.com/superjcd/shopservice/service"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.ShopServiceServer, error) {

	wire.Build(
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
