package dataset_db

import (
	"dataset/internal/infrastructure/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	// dbInstance is global dataset_db instance
	dbInstance *gorm.DB

	// once is used to ensure that the dataset_db instance is initialized only once
	once sync.Once
)

// InitializeDB initializes the dataset_db instance, which is used to connect to the database,create tables, etc.
func InitializeDB() {
	//dataset_db := GetDB()
	// create table if not exists
	//_ = dataset_db.AutoMigrate(&po.DatasetPO{})
	//_ = dataset_db.AutoMigrate(&po.MetricPO{})
	//_ = dataset_db.AutoMigrate(&po.DimensionPO{})
	//_ = dataset_db.AutoMigrate(&po.DataSourcePO{})
	//_ = dataset_db.AutoMigrate(&po.DataSourceFilePO{})
	//_ = dataset_db.AutoMigrate(&po.DataSourceMySQLPO{})
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
