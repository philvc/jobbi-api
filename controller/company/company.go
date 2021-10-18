package company_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type CompanyController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) CompanyController {
	return CompanyController{
		usecase: usecase,
	}
}

// swagger:operation GET /searches/{searchId}/companies company GetCompaniesBySearchId
// type id struct
// Get Companies by searchId.
// Return company
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/CompanyDTO"
//       400:
//         description: Bad Request

func (controller CompanyController) GetCompaniesBySearchId(c *gin.Context) {
	searchId := c.Params.ByName("searchId")

	Companies, error := controller.usecase.CompanyUsecase.GetCompaniesBySearchId(searchId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Companies)
}

func (controller CompanyController) GetCompanyById(c *gin.Context) {

	companyId := c.Params.ByName("companyId")

	Company, error := controller.usecase.CompanyUsecase.GetCompanyById(companyId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Company)
}

func (controller CompanyController) AddCompany(c *gin.Context) {

	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	var Company contract.CompanyDTO

	if err := c.BindJSON(&Company); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	searchDTO, err := controller.usecase.SearchUsecase.GetSearchById(searchId)
	userDTO, err := controller.usecase.UserUsecase.GetUserBySub(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	Company.UserID = userDTO.Id
	Company.SearchID = searchDTO.Id

	CompanyDTO, err := controller.usecase.CompanyUsecase.AddCompany(Company)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, CompanyDTO)
}

func (controller CompanyController) ModifyCompany(c *gin.Context) {
	var Company contract.CompanyDTO

	if err := c.BindJSON(&Company); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	CompanyDTO, error := controller.usecase.CompanyUsecase.ModifyCompany(Company)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, CompanyDTO)
}
