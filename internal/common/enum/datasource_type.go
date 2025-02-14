package enum

type DataSourceType string

const (
	// File is a type of data source
	File DataSourceType = "file"

	// MySQL is a type of data source
	MySQL DataSourceType = "mysql"

	// ExcelDataSource is a type of data source
	ExcelDataSource DataSourceType = "excel"

	// CSVDataSource is a type of data source
	CSVDataSource DataSourceType = "csv"
)

func GetDataSourceType(val string) DataSourceType {
	return DataSourceType(val)
}
