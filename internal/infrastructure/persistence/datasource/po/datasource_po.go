package po

import (
	"dataset/internal/infrastructure/persistence/db"
)

// DataSourcePO the t
type DataSourcePO struct {
	// Model base field
	db.Model

	// Description  to describe the data source
	Description string `gorm:"column:description;type:VARCHAR(255);not null"`

	// Name the name of the data sources
	Name string `gorm:"column:name;type:VARCHAR(255);not null;unique"`
}

func (*DataSourcePO) TableName() string {
	return "t_datasource"
}
