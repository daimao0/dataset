package dto

import "dataset/internal/common/enum"

// DataSourceMySQLDTO is the data source mysql
type DataSourceMySQLDTO struct {

	// BaseDataSourceDTO is the base data source dto
	BaseDataSourceDTO

	// Address is the dns of mysql
	Address string

	// Database is the database of mysql
	Database string

	// User is the user of mysql
	Username string

	// Password is the password of mysql
	Password string
}

func (d *DataSourceMySQLDTO) DataSourceDTOType() enum.DataSourceType {
	return enum.MySQL
}
