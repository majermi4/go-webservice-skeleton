package db

import (
	"MyWebService/config"
	"context"
	"database/sql"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config, lc fx.Lifecycle) (*gorm.DB, *sql.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DbDsn), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Ping the database to ensure connection is established
			return sqlDB.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return sqlDB.Close()
		},
	})

	return db, sqlDB, nil
}
