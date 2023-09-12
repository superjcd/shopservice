package server

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/HooYa-Bigdata/shopservice/genproto/v1/gw"
	"github.com/HooYa-Bigdata/shopservice/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RunHttpServer Run http server
func RunHttpServer(cfg *config.Config) {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := v1.RegisterShopServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		cfg.Grpc.Port,
		opts,
	); err != nil {
		panic("register service handler failed.[ERROR]=>" + err.Error())
	}

	httpServer := &http.Server{
		Addr:    cfg.Http.Port,
		Handler: mux,
	}
	fmt.Println("Listening http server on port" + cfg.Http.Port)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			fmt.Println("listen http server failed.[ERROR]=>" + err.Error())
		}
	}()

	cfg.Http.Server = httpServer

}
