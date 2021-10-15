package invoice_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/usecase"
)

type InvoiceController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) InvoiceController {
	return InvoiceController{
		usecase: usecase,
	}
}

func (controller InvoiceController) GetInvoices(c *gin.Context) {

	organisationId := c.Params.ByName("organisationId")
	invoices, error := controller.usecase.InvoiceUsecase.GetInvoicesByOrganisationId(organisationId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, invoices)
}

func (controller InvoiceController) GetInvoiceById(c *gin.Context) {

	invoiceId := c.Params.ByName("invoiceId")
	invoice, error := controller.usecase.InvoiceUsecase.GetInvoiceById(invoiceId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, invoice)
}
