package impl

import (
	"dataset/internal/application/cmd"
	"dataset/internal/application/convert"
	"dataset/internal/application/dto"
	"dataset/internal/application/factory"
	"dataset/internal/domain/physical_table/service"
)

// PhysicalTableAppImpl is the application interface implementation
type PhysicalTableAppImpl struct {
	physicalTableService service.PhysicalTableService
}

func NewPhysicalTableAppImpl() *PhysicalTableAppImpl {
	return &PhysicalTableAppImpl{
		//physicalTableService: impl.NewBasePhysicalTableServiceImpl(),
	}
}

// CreateByData creates a physical table by data
func (p *PhysicalTableAppImpl) CreateByData(cmd *cmd.PhysicalTableCreateByDataCmd) (*dto.PhysicalTableDTO, error) {
	physicalTable := factory.CreatePhysicalTableFormData(cmd)
	tableDTO := convert.PhysicalTableTOPhysicalTableDTO(physicalTable)
	return tableDTO, nil
}

// CreateBySQL creates a physical table by SQL
func (p *PhysicalTableAppImpl) CreateBySQL(cmd *cmd.PhysicalTableCreateBySQLCmd) (*dto.PhysicalTableDTO, error) {

	return nil, nil
}
