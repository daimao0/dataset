package model

import "dataset/internal/common/enum"

// FileDataSource is a data source that reads data from a file.
type FileDataSource struct {

	// Path is the path to the file.(oss or local file)
	Path string

	// FileType is the type of the file.
	FileType enum.FileType
}

// ReadData reads the data from the file
func (f *FileDataSource) ReadData() ([]map[string]any, error) {
	return nil, nil
}

func (f *FileDataSource) readCSV() ([]map[string]any, error) {
	return nil, nil
}
