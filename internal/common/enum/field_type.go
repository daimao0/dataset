package enum

type FieldType string

const (
	// MeasureType is a measure
	MeasureType FieldType = "measure"
	// MetricType is a metric
	MetricType FieldType = "metric"
	// DimensionType is a dimension
	DimensionType FieldType = "dimension"
	// DriverDimensionType is a driver dimension
	DriverDimensionType FieldType = "driver_dimension"
)

func GetFieldType(val string) FieldType {
	switch val {
	case "measure":
		return MeasureType
	case "metric":
		return MetricType
	case "dimension":
		return DimensionType
	case "driver_dimension":
		return DriverDimensionType
	default:
		return DimensionType
	}
}
