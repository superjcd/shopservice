package store

import (
	v1 "github.com/HooYa-Bigdata/shopservice/genproto/v1"
	"gorm.io/gorm"
)

// gorm 模型定义

type Shop struct {
	gorm.Model
	Name       string `json:"name" gorm:"column:name"`
	BrandNames string `json:"password" gorm:"column:brand_names"`
	Country    string `json:"country" gorm:"column:country"`
	Tags       string `json:"tags" gorm:"column:tags"`
}

type ShopList struct {
	TotalCount int64   `json:"totalCount"`
	Items      []*Shop `json:"items"`
}

func (sl *ShopList) ConvertToListShopResponse(msg string, status v1.Status) v1.ListShopResponse {
	shops := make([]*v1.Shop, 8)

	for _, shop := range sl.Items {
		shops = append(shops, &v1.Shop{
			Name:       shop.Name,
			BrandNames: shop.BrandNames,
			Country:    shop.Country,
			Tags:       shop.Tags,
		})
	}

	return v1.ListShopResponse{
		Msg:    msg,
		Status: status,
		Shops:  shops,
	}
}

func MigrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(Shop{}); err != nil {
		return err
	}
	return nil
}
