package utils

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	response "github.com/vicradon/internpulse/stage4/pkg"
)

func HandleValidationError(c *gin.Context, err error) {
	log.Println(err)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var missingFields []string
		for _, fieldErr := range validationErrors {
			missingFields = append(missingFields, fieldErr.Field())
		}
		response.Error(c, http.StatusBadRequest, "missing or invalid fields: "+strings.Join(missingFields, ", "))
	} else {
		response.Error(c, http.StatusBadRequest, "error reading request body: "+err.Error())
	}
}
