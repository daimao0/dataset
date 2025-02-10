package model

// DataSource is the interface for data source
type DataSource interface {

	// ReadData reads data from data source
	ReadData() (map[string]interface{}, error)
}
