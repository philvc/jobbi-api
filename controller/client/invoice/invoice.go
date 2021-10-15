package client_invoice_controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type ClientInvoiceController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) ClientInvoiceController {
	return ClientInvoiceController{
		usecase: usecase,
	}
}

func (controller ClientInvoiceController) GetInvoices(c *gin.Context) {

	clientId := c.Params.ByName("clientId")
	invoices, error := controller.usecase.InvoiceUsecase.GetInvoicesByClientId(clientId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, invoices)
}

func (controller ClientInvoiceController) GetInvoiceById(c *gin.Context) {

	invoiceId := c.Params.ByName("invoiceId")
	invoice, error := controller.usecase.InvoiceUsecase.GetInvoiceById(invoiceId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, invoice)
}

func (controller ClientInvoiceController) AddInvoice(c *gin.Context) {

	var invoice contract.InvoiceDTO

	if err := c.BindJSON(&invoice); err != nil {
		log.Default().Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	invoiceDTO, error := controller.usecase.InvoiceUsecase.AddInvoice(invoice)

	if error != nil {
		log.Default().Println(error.Error())
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, invoiceDTO)
}

func (controller ClientInvoiceController) ModifyInvoice(c *gin.Context) {

	var invoice contract.InvoiceDTO

	if err := c.BindJSON(&invoice); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	invoiceDTO, error := controller.usecase.InvoiceUsecase.ModifyInvoice(invoice)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, invoiceDTO)
}
