package controller

import (
	"dataset/internal/adapter/http/request"
	"dataset/internal/application/app"
	"dataset/internal/application/app/impl"
	"dataset/internal/application/cmd"
	"dataset/internal/common/easytool/resp"
	"dataset/internal/common/easytool/util/csv_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DatasetController struct {
	datasetApp app.DatasetApplication
}

func NewDatasetController() *DatasetController {
	return &DatasetController{
		datasetApp: impl.NewDatasetAppImpl(),
	}
}

// GetDataset gets dataset by id
func (d *DatasetController) GetDataset(c *gin.Context) {
	c.JSON(http.StatusOK, resp.Success("a"))
}

// CreateDatasetByJSON creates dataset by json array
func (d *DatasetController) CreateDatasetByJSON(c *gin.Context) {
	datasetCreateJSON := &request.DatasetCreateJSON{}
	err := c.ShouldBindJSON(datasetCreateJSON)
	if err != nil {
		c.JSON(http.StatusOK, resp.InvalidParam("data can not be empty"))
		return
	}
	dto, err := d.datasetApp.CreateDatasetByData(&cmd.DatasetCreateByDataCmd{Name: datasetCreateJSON.Name, Data: datasetCreateJSON.Data})
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(dto))
}

// CreateDatasetByFile creates dataset by file
func (d *DatasetController) CreateDatasetByFile(c *gin.Context) {
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
	data, err := d.datasetApp.CreateDatasetByData(&cmd.DatasetCreateByDataCmd{Name: file.Filename, Data: records})
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(data))
}
