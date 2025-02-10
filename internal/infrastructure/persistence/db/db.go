package db

import (
	"dataset/internal/infrastructure/config"
	"dataset/internal/infrastructure/persistence/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	// dbInstance is global db instance
	dbInstance *gorm.DB

	// once is used to ensure that the db instance is initialized only once
	once sync.Once
)

// InitializeDB initializes the db instance, which is used to connect to the database,create tables, etc.
func InitializeDB() {
	db := GetDB()
	// create table if not exists
	_ = db.AutoMigrate(&po.DatasetPO{})
	_ = db.AutoMigrate(&po.MeasurePO{})
	_ = db.AutoMigrate(&po.MetricPO{})
	_ = db.AutoMigrate(&po.DimensionPO{})
	_ = db.AutoMigrate(&po.DriverDimensionPO{})
}

func GetDB() *gorm.DB {
	once.Do(func() {
		dbInstance = initDB()
	})
	return dbInstance
}

func initDB() *gorm.DB {
	appConfig := config.AppConfig
	db, err := gorm.Open(mysql.Open(appConfig.Database.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(appConfig.Database.MaxOpenConn)
	sqlDB.SetMaxIdleConns(appConfig.Database.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(appConfig.Database.MaxLifetime) * time.Second)
	return db
}
