package service

import (
	"context"

	"github.com/superjcd/shopservice/config"
	v1 "github.com/superjcd/shopservice/genproto/v1"
	"github.com/superjcd/shopservice/pkg/database"
	"github.com/superjcd/shopservice/service/store"
	"github.com/superjcd/shopservice/service/store/sql"
	"gorm.io/gorm"
)

var _DB *gorm.DB

// Server Server struct
type Server struct {
	v1.UnimplementedShopServiceServer
	datastore store.Factory
	client    v1.ShopServiceClient
	conf      *config.Config
}

// NewServer New service grpc server
func NewServer(conf *config.Config, client v1.ShopServiceClient) (v1.ShopServiceServer, error) {
	_DB = database.MustPreParePostgresqlDb(&conf.Pg)
	factory, err := sql.NewSqlStoreFactory(_DB)
	if err != nil {
		return nil, err
	}

	server := &Server{
		client:    client,
		datastore: factory,
		conf:      conf,
	}

	return server, nil
}

func (s *Server) CreateShop(ctx context.Context, rq *v1.CreateShopRequest) (*v1.CreateShopResponse, error) {
	if err := s.datastore.Shops().Create(ctx, rq); err != nil {
		return &v1.CreateShopResponse{Msg: "创建失败", Status: v1.Status_failure, Shop: nil}, err
	}
	return &v1.CreateShopResponse{
		Msg:    "创建成功",
		Status: v1.Status_success,
		Shop:   &v1.Shop{Name: rq.Name, BrandNames: rq.BrandNames, Country: rq.Country, Tags: rq.Tags},
	}, nil
}

func (s *Server) ListShop(ctx context.Context, rq *v1.ListShopRequest) (*v1.ListShopResponse, error) {
	groups, err := s.datastore.Shops().List(ctx, rq)
	if err != nil {
		return &v1.ListShopResponse{Msg: "获取列表失败", Status: v1.Status_failure}, err
	}

	resp := groups.ConvertToListShopResponse("成功获取列表", v1.Status_success)

	return &resp, nil
}

func (s *Server) UpdateShop(ctx context.Context, rq *v1.UpdateShopRequest) (*v1.UpdateShopResponse, error) {
	if err := s.datastore.Shops().Update(ctx, rq); err != nil {
		return &v1.UpdateShopResponse{Msg: "失败", Status: v1.Status_success}, err
	}

	return &v1.UpdateShopResponse{
		Msg:    "更新成功",
		Status: v1.Status_success,
	}, nil

}

func (s *Server) DeleteShop(ctx context.Context, rq *v1.DeleteShopRequest) (*v1.DeleteShopResponse, error) {
	if err := s.datastore.Shops().Delete(ctx, rq); err != nil {
		return &v1.DeleteShopResponse{Msg: "删除失败", Status: v1.Status_failure}, err
	}

	return &v1.DeleteShopResponse{Msg: "删除成功", Status: v1.Status_success}, nil
}
