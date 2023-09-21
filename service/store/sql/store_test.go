package sql

import (
	"context"
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	v1 "github.com/superjcd/shopservice/genproto/v1"
	"github.com/superjcd/shopservice/service/store"
	"gorm.io/gorm"
)

var dbFile = "fake.db"

type FakeStoreTestSuite struct {
	suite.Suite
	Dbfile      string
	FakeFactory store.Factory
}

func (suite *FakeStoreTestSuite) SetupSuite() {
	file, err := os.Create(dbFile)
	assert.Nil(suite.T(), err)
	defer file.Close()

	suite.Dbfile = dbFile
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	assert.Nil(suite.T(), err)

	factory, err := NewSqlStoreFactory(db)
	assert.Nil(suite.T(), err)
	suite.FakeFactory = factory
}

func (suite *FakeStoreTestSuite) TearDownSuite() {
	var err error
	err = suite.FakeFactory.Close()
	assert.Nil(suite.T(), err)
	err = os.Remove(dbFile)
	assert.Nil(suite.T(), err)
}

func (suite *FakeStoreTestSuite) TestCreateShop() {
	newShop := &v1.CreateShopRequest{
		Name:       "apple",
		BrandNames: "apple;apple tv",
		Country:    "US",
		Tags:       "amazon",
	}

	err := suite.FakeFactory.Shops().Create(context.Background(), newShop)
	assert.Nil(suite.T(), err)
}

func (suite *FakeStoreTestSuite) TestListShops() {
	request := &v1.ListShopRequest{
		Part:   "app",
		Offset: 0, // 必须要用到part,
		Limit:  10,
	}
	shopList, err := suite.FakeFactory.Shops().List(context.Background(), request)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(shopList.Items))
}

func (suite *FakeStoreTestSuite) TestUpdateShop() {
	newShop := &v1.UpdateShopRequest{
		Name:       "apple",
		BrandNames: "apple;apple tv",
		Country:    "CA",
		Tags:       "amazon",
	}
	err := suite.FakeFactory.Shops().Update(context.Background(), newShop)
	assert.Nil(suite.T(), err)

	// 确认一下修改是否成功
	request := &v1.ListShopRequest{
		Part:   "app",
		Offset: 0, // 必须要用到part,
		Limit:  10,
	}
	shopList, err := suite.FakeFactory.Shops().List(context.Background(), request)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "CA", shopList.Items[0].Country)

}

func (suite *FakeStoreTestSuite) TestZDeleteShop() { // 添加Z的原因是希望能够靠后运行这个测试
	target := &v1.DeleteShopRequest{
		Name: "apple",
	}
	err := suite.FakeFactory.Shops().Delete(context.Background(), target)
	assert.Nil(suite.T(), err)

	// 确认一下修改是否成功
	request := &v1.ListShopRequest{
		Part:   "app",
		Offset: 0, // 必须要用到part,
		Limit:  10,
	}
	shopList, err := suite.FakeFactory.Shops().List(context.Background(), request)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, len(shopList.Items))

}

func TestFakeStoreSuite(t *testing.T) {
	suite.Run(t, new(FakeStoreTestSuite))
}
