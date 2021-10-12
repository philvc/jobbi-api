package invoice_export_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/usecase"
)

type ExportController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) ExportController {
	return ExportController{
		usecase: usecase,
	}
}

func (controller ExportController) GetExport(c *gin.Context) {

	// Gets invoice
	invoiceId := c.Params.ByName("invoiceId")

	encoded, err := controller.usecase.ExportUsecase.GetExportByInvoiceId(invoiceId)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, encoded)
}
