package po

type DriverDimensionPO struct {
	// Model default columns
	Model

	// Name of Dimension
	Name string `gorm:"column:name;type:VARCHAR(255);not null"`

	// Label of Dimension
	Label string `gorm:"column:label;type:VARCHAR(255);not null;default:''"`

	// DataType of Dimension
	DataType string `gorm:"column:data_type;type:VARCHAR(255);not null"`

	// Description of dataset
	Description string `gorm:"column:description;type:VARCHAR(255);not null"`
}

func (*DriverDimensionPO) TableName() string {
	return "t_driver_dimension"
}
