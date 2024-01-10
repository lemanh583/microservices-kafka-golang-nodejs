package database

import (
	"fmt"
	"time"
	"user-services/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getStringConnect() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.DatabaseUsername,
		config.Cfg.DatabasePassword,
		config.Cfg.DatabaseHost,
		config.Cfg.DatabasePort,
		config.Cfg.DatabaseName,
	)
}

func ConnectToDatabase() (*gorm.DB, error) {
	var configDB *mysql.Config = &mysql.Config{
		DSN: getStringConnect(),
	}
	fmt.Println(getStringConnect())
	db, err := gorm.Open(mysql.New(*configDB), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = setConnectionPool(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
