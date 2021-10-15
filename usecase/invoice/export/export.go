package invoice_export_usecase

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/philvc/jobbi-api/repository"
	"github.com/signintech/gopdf"
)

type ExportUsecase struct {
	repository repository.Repository
}

// Returns an instance of a invoice use-case
func Default(repository repository.Repository) ExportUsecase {
	return ExportUsecase{
		repository: repository,
	}
}

func (usecase ExportUsecase) GetExportByInvoiceId(invoiceId string) (string, error) {

	// Get Invoice
	invoice, err := usecase.repository.InvoiceRepository.GetInvoiceById(invoiceId)
	if err != nil {
		return "", err
	}

	// Get Client
	client, err := usecase.repository.ClientRepository.GetClientById(invoice.ClientId)
	if err != nil {
		return "", err
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err = addFonts(&pdf)
	if err != nil {
		return "", errors.New("error/adding-fonts")
	}

	// Adds the header image
	rect := gopdf.Rect{
		W: 595,
		H: 63,
	}
	pdf.Image("assets/images/header.jpg", 0, 0, &rect)

	// Provider - Billing information
	addText(&pdf, "Bill from", "Poppins-SemiBold", 12, 36, 97)
	addText(&pdf, "Nightborn Agency", "Poppins-SemiBold", 10, 36, 117)
	addText(&pdf, "Rue Notre-Seigneur 11", "Poppins-Regular", 10, 36, 137)
	addText(&pdf, "1000, Brussels", "Poppins-Regular", 10, 36, 157)
	addText(&pdf, "BTW BE123.344.344.342", "Poppins-Regular", 10, 36, 177)
	addText(&pdf, "IBAN BE123.344.344.342", "Poppins-Regular", 10, 36, 197)
	addText(&pdf, "SWIFT (BIC) ", "Poppins-Regular", 10, 36, 217)

	// Provider - Contact information
	addText(&pdf, "Phone", "Poppins-SemiBold", 10, 36, 247)
	addText(&pdf, "+32 494 90 36 65", "Poppins-Regular", 10, 86, 247)

	addText(&pdf, "E-mail", "Poppins-SemiBold", 10, 36, 267)
	addText(&pdf, "invoices@nightborn.be", "Poppins-Regular", 10, 86, 267)

	// Date section
	addText(&pdf, "Date", "Poppins-SemiBold", 12, 349, 97)
	addText(&pdf, invoice.Date.Format("02/01/2006"), "Poppins-Regular", 10, 349, 117)

	// Invoice number section
	addText(&pdf, "Invoice #", "Poppins-SemiBold", 12, 349, 162)
	addText(&pdf, invoice.Number, "Poppins-Regular", 10, 349, 182)

	// Client - Billing information
	addText(&pdf, "Bill to", "Poppins-SemiBold", 12, 36, 312)
	addText(&pdf, client.Name, "Poppins-SemiBold", 10, 36, 332)
	addText(&pdf, client.Address, "Poppins-Regular", 10, 36, 352)
	addText(&pdf, client.VATNumber, "Poppins-Regular", 10, 36, 372)

	// Service information
	addText(&pdf, "For", "Poppins-SemiBold", 12, 349, 312)
	addText(&pdf, invoice.Description, "Poppins-Regular", 10, 349, 332)

	// Adds the table header
	rect = gopdf.Rect{
		W: 595,
		H: 27,
	}
	pdf.Image("assets/images/table-header.jpg", 0, 427, &rect)

	// Table
	addText(&pdf, invoice.Description, "Poppins-SemiBold", 10, 36, 477)
	addText(&pdf, fmt.Sprintf("%v", invoice.Amount), "Poppins-Regular", 10, 520, 477)

	// Adds a line
	rect = gopdf.Rect{
		W: 210,
		H: 1,
	}
	pdf.Image("assets/images/line.jpg", 349, 488, &rect)

	// Sub-total
	addText(&pdf, "Subtotal", "Poppins-SemiBold", 10, 349, 507)
	addText(&pdf, fmt.Sprintf("%v", invoice.SubTotal), "Poppins-Regular", 10, 520, 507)

	// VAT
	addText(&pdf, "VAT", "Poppins-SemiBold", 10, 349, 527)
	addText(&pdf, fmt.Sprintf("%v", invoice.VAT), "Poppins-Regular", 10, 520, 527)

	// Adds a line
	rect = gopdf.Rect{
		W: 210,
		H: 1,
	}
	pdf.Image("assets/images/line.jpg", 349, 538, &rect)

	// Total
	addText(&pdf, "TOTAL", "Poppins-Bold", 10, 349, 557)
	addText(&pdf, fmt.Sprintf("%v", invoice.Total), "Poppins-Regular", 10, 520, 557)

	// Adds the footer
	rect = gopdf.Rect{
		W: 595,
		H: 40,
	}
	pdf.Image("assets/images/footer.jpg", 0, 803, &rect)

	// Creates the PDF
	filePath := "./tmp/invoice-" + invoice.Number + ".pdf"
	pdf.WritePdf(filePath)

	// Open file on disk.
	f, _ := os.Open(filePath)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Deletes the TMP file from the server
	os.Remove(filePath)

	return encoded, nil
}

func addText(pdf *gopdf.GoPdf, text string, font string, size int, x float64, y float64) {
	pdf.SetFont(font, "", size)
	pdf.SetX(x)
	pdf.SetY(y)
	pdf.Text(text)
}

func addFonts(pdf *gopdf.GoPdf) error {
	err := pdf.AddTTFFont("Poppins-Black", "assets/fonts/Poppins-Black.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-ExtraBold", "assets/fonts/Poppins-ExtraBold.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-Bold", "assets/fonts/Poppins-Bold.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-SemiBold", "assets/fonts/Poppins-SemiBold.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-Regular", "assets/fonts/Poppins-Regular.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-Medium", "assets/fonts/Poppins-Medium.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-Light", "assets/fonts/Poppins-Light.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-ExtraLight", "assets/fonts/Poppins-ExtraLight.ttf")
	if err != nil {
		return err
	}

	err = pdf.AddTTFFont("Poppins-Thin", "assets/fonts/Poppins-Thin.ttf")
	if err != nil {
		return err
	}

	return nil
}
