package controller

import (
	"dataset/internal/application/app"
	"dataset/internal/application/app/impl"
	"dataset/internal/application/cmd"
	"dataset/internal/common/easytool/resp"
	"dataset/internal/common/easytool/util/csv_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PhysicalTableController struct {
	physicalTableApp app.PhysicalTableApp
}

func NewPhysicalTableController() *PhysicalTableController {
	return &PhysicalTableController{
		physicalTableApp: impl.NewPhysicalTableAppImpl(),
	}
}

// CreatePhysicalByFile creates dataset by file
func (p *PhysicalTableController) CreatePhysicalByFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, resp.InvalidParam("file can not be empty"))
		return
	}
	open, err := file.Open()
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	records, err := csv_util.ReadCSVByReader(open)
	dto, err := p.physicalTableApp.CreateByData(&cmd.PhysicalTableCreateByDataCmd{Name: file.Filename, Data: records})
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(dto))
}
