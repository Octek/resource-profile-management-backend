package projects

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Routes Exports all routes handled by this service
func Routes(router *gin.Engine, projectSvc ProjectService) {

}
