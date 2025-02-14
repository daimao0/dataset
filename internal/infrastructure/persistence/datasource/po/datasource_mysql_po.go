package po

// DataSourceMySQLPO the table is mysql type data source
type DataSourceMySQLPO struct {

	// DataSourcePO is the base info of data source
	DataSourcePO

	// Address is the dns of mysql
	Address string

	// Database is the database of mysql
	Database string

	// User is the user of mysql
	Username string

	// Password is the password of mysql
	Password string
}

func (*DataSourceMySQLPO) TableName() string {
	return "t_datasource_mysql"
}
