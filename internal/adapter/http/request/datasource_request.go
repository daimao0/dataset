package request

// MySQLDataSourceRequest is the request of MySQL data source.
type MySQLDataSourceRequest struct {

	// Id is the id of data source.
	// If Id is empty, it will be created. If Id is not empty, it will be updated.
	Id string

	// Name is the name of MySQL data source. It created by user.
	Name string `json:"name"`

	// Description is the description of MySQL data source.
	Description string `json:"description"`

	// Address is the address of MySQL.
	Address string `json:"address"`

	// User is the user of MySQL.
	Username string `json:"username"`

	// Password is the password of MySQL.
	Password string `json:"password"`

	// Database is the database of MySQL.
	Database string `json:"database"`
}
