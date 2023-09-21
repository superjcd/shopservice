package sql

import (
	"fmt"
	"sync"

	"github.com/superjcd/shopservice/service/store"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Shops() store.ShopStore {
	return &shops{db: ds.db}
}

var (
	sqlFactory store.Factory
	once       sync.Once
)

func NewSqlStoreFactory(db *gorm.DB) (store.Factory, error) {
	if db == nil && sqlFactory == nil {
		return nil, fmt.Errorf("failed to get pg store fatory")
	}
	once.Do(func() {
		store.MigrateDatabase(db)
		sqlFactory = &datastore{db: db}
	})

	return sqlFactory, nil
}

func (ds *datastore) Close() error {
	db, _ := ds.db.DB()

	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
