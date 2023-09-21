package sql

import (
	"context"
	"fmt"

	v1 "github.com/superjcd/shopservice/genproto/v1"
	"github.com/superjcd/shopservice/service/store"
	"gorm.io/gorm"
)

type shops struct {
	db *gorm.DB
}

var _ store.ShopStore = (*shops)(nil)

func (s *shops) Create(ctx context.Context, rq *v1.CreateShopRequest) error {

	group := store.Shop{
		Name:       rq.Name,
		BrandNames: rq.BrandNames,
		Country:    rq.Country,
		Tags:       rq.Tags,
	}

	return s.db.Create(&group).Error // 我只存储了用户， 但没有处理和用户group有关的逻辑
}

func (s *shops) List(ctx context.Context, rq *v1.ListShopRequest) (*store.ShopList, error) {
	result := &store.ShopList{}

	var where_clause string
	if rq.Part == "" {
		where_clause = "1 = 1"
	} else {
		where_clause = fmt.Sprintf("name like '%%%s%%'", rq.Part)
	}

	d := s.db.Where(where_clause).
		Offset(int(rq.Offset)).
		Limit(int(rq.Limit)).
		Find(&result.Items).
		Offset(-1).
		Limit(-1).
		Count(&result.TotalCount)

	return result, d.Error
}

func (s *shops) Update(ctx context.Context, rq *v1.UpdateShopRequest) error {
	shop := store.Shop{}
	if err := s.db.Where("name = ?", rq.Name).First(&shop).Error; err != nil {
		return err
	}
	shop.BrandNames = rq.BrandNames
	shop.Country = rq.Country
	shop.Tags = rq.Tags
	return s.db.Save(&shop).Error
}

func (s *shops) Delete(ctx context.Context, rq *v1.DeleteShopRequest) error {
	return s.db.Where("name = ?", rq.Name).Delete(&store.Shop{}).Error
}
