package pgdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGClient struct {
	DB *gorm.DB
}

func Connect(dsn string) (*PGClient, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return &PGClient{db}, nil
}

func (c *PGClient) Close() error {

	sqlDB, err := c.DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()

}

func (c *PGClient) Ping() error {
	sqlDB, err := c.DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Ping()
}
