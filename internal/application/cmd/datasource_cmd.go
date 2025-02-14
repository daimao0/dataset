package cmd

import "dataset/internal/common/enum"

// DataSourceCmd is a cmd for edit a data source.
type DataSourceCmd interface {

	// DataSourceCmdType is the type of data source.
	DataSourceCmdType() enum.DataSourceType
}

// BaseDataSourceCmd is a cmd for edit a base data source.
type BaseDataSourceCmd struct {

	// Id is the id of data source.
	// If Id is empty, it will be created. If Id is not empty, it will be updated.
	Id string

	// Name is the name of data source.
	Name string

	// Description is the description of data source.
	Description string
}

// MySQLDataSourceCmd is a cmd for edit a MySQL data source.
type MySQLDataSourceCmd struct {

	// BaseDataSourceCmd is the base data source cmd.
	BaseDataSourceCmd

	// Address is the address of MySQL.
	Address string

	// User is the user of MySQL.
	Username string

	// Password is the password of MySQL.
	Password string

	// Database is the database of MySQL.
	Database string
}

func (m *MySQLDataSourceCmd) DataSourceCmdType() enum.DataSourceType {
	return enum.MySQL
}
