package po

import "dataset/internal/infrastructure/persistence/dataset_db"

// DataSourceMySQLPO the table is mysql type data source
type DataSourceMySQLPO struct {

	// Model is database info
	dataset_db.Model

	// Address is the dns of mysql
	Address string `gorm:"column:address;type:varchar(255);not null;"`

	// Database is the database of mysql
	Database string `gorm:"column:database;type:varchar(255);not null;"`

	// User is the user of mysql
	Username string `gorm:"column:username;type:varchar(255);not null;"`

	// Password is the password of mysql
	Password string `gorm:"column:password;type:varchar(255);not null;"`
}

func (*DataSourceMySQLPO) TableName() string {
	return "t_datasource_mysql"
}
