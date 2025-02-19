package po

import "dataset/internal/infrastructure/persistence/dataset_db"

type MetricPO struct {
	// Model default columns
	dataset_db.Model

	// Name of Dimension
	Name string `gorm:"column:name;type:VARCHAR(255);not null"`

	// Label of Dimension
	Label string `gorm:"column:label;type:VARCHAR(255);not null;default:''"`

	// DataType of Dimension
	DataType string `gorm:"column:data_type;type:VARCHAR(255);not null"`

	// Description of dataset
	Description string `gorm:"column:description;type:VARCHAR(255);not null"`
}

func (*MetricPO) TableName() string {
	return "t_metric"
}
