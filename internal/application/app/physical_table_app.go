package app

import (
	"dataset/internal/application/cmd"
	"dataset/internal/application/dto"
)

// PhysicalTableApp is physical table  application interface
type PhysicalTableApp interface {

	// CreateByData create a physical table by data
	CreateByData(cmd *cmd.PhysicalTableCreateByDataCmd) (*dto.PhysicalTableDTO, error)

	// CreateBySQL create a physical table by sql
	CreateBySQL(cmd *cmd.PhysicalTableCreateBySQLCmd) (*dto.PhysicalTableDTO, error)
}
