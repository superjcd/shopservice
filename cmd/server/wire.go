//go:build wireinject
// +build wireinject

package server

import (
	"github.com/HooYa-Bigdata/shopservice/config"
	v1 "github.com/HooYa-Bigdata/shopservice/genproto/v1"
	"github.com/HooYa-Bigdata/shopservice/service"
	"github.com/google/wire"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.ShopServiceServer, error) {

	wire.Build(
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
