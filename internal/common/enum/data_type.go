package enum

import (
	"dataset/internal/common/easytool/util/str_util"
	"strconv"
	"time"
)

// DataType is the data type for field
type DataType string

const (
	// Number is a number,contains integer and float
	Number DataType = "number"

	// String is a string
	String DataType = "string"

	// DateTime is a date time
	DateTime DataType = "date_time"

	// Enum is a enum
	Enum DataType = "enum"
)

// GetDataType get data type enum
func GetDataType(val string) DataType {

	switch val {
	case "number":
		return Number
	case "string":
		return String
	case "date_time":
		return DateTime
	case "enum":
		return Enum
	default:
		return String
	}
}

// ParseToDataType parse any value to data type
func ParseToDataType(obj any) DataType {
	switch v := obj.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return Number
	case string:
		return parseStringToDataType(v)
	case time.Time:
		return DateTime
	default:
		return String
	}
}

// parseStringToDataType parse string value to data type
func parseStringToDataType(val string) DataType {
	if _, err := strconv.ParseFloat(val, 64); err == nil {
		return Number
	}
	if _, err := str_util.ParseDatetime(val); err == nil {
		return DateTime
	}
	return String
}
