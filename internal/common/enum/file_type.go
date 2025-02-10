package enum

import "errors"

type FileType string

const (
	CSV   FileType = "csv_util"
	EXCEL FileType = "excel"
)

func GetFileType(val string) (FileType, error) {
	switch val {
	case "csv_util":
		return CSV, nil
	case "excel":
		return EXCEL, nil
	}
	return "", errors.New("no file type")
}
