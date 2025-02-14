package convert

import (
	"dataset/internal/application/dto"
	"dataset/internal/domain/physical_table/model"
)

func PhysicalTableTOPhysicalTableDTO(table *model.PhysicalTable) *dto.PhysicalTableDTO {

	return &dto.PhysicalTableDTO{
		CreatedAt:   table.CreatedAt,
		Description: table.Description,
		Fields:      FieldSToFieldDTOS(table.Fields),
		Id:          table.Id,
		Name:        table.Name,
		OriginName:  table.OriginName,
		UpdatedAt:   table.UpdatedAt,
	}
}
