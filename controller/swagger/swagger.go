package swagger_controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SwaggerController struct{}

func InitialiseSwaggerController() SwaggerController {
	return SwaggerController{}
}

func (controller SwaggerController) GetSwagger(c *gin.Context) {
	// Read content
	bytes, err := os.ReadFile("swagger.json")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Parse bytes to json
	var result map[string]interface{}
	err = json.Unmarshal([]byte(bytes), &result)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
