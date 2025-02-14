package model

import "dataset/internal/common/enum"

// DataSourceMySQL is mysql data source
type DataSourceMySQL struct {
	// BaseDataSource is the base data source.
	BaseDataSource

	// Address is the address of MySQL.
	Address string

	// Username is the user of MySQL.
	Username string

	// Password is the password of MySQL.
	Password string

	// Database is the database of MySQL.
	Database string
}

func (d *DataSourceMySQL) DataSourceType() enum.DataSourceType {
	return enum.MySQL
}
