package po

// DataSourceFilePO the table is file type data source
type DataSourceFilePO struct {

	// DataSourcePO is the base info of data source
	DataSourcePO

	// FilePath is the path of file
	FilePath string `gorm:"column:file_path;type:VARCHAR(255);not null"`

	// FileName is the name of file
	FileName string `gorm:"column:file_name;type:VARCHAR(255);not null"`
}

func (*DataSourceFilePO) TableName() string {
	return "t_datasource_file"
}
