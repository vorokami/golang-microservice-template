package http

import (
	"encoding/json"
	"fmt"
	"golang-microservice-template/internal/model"
	"golang-microservice-template/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type templateRoutes struct {
	uc     usecase.TemplateMethods
	logger *zap.Logger
}

func newTemplateRoutes(handler *gin.RouterGroup, uc usecase.TemplateMethods, logger *zap.Logger) {
	r := &templateRoutes{uc, logger}

	handler.GET("/get", r.getMethodTemplate)
	handler.POST("/post", r.postMethodTemplate)
}

func (r *templateRoutes) getMethodTemplate(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	// if there is a typo in some fields, we will get an error
	decoder.DisallowUnknownFields()

	var request model.GetMethodTemplateRequest
	err := decoder.Decode(&request)
	if err != nil {
		r.logger.Error("Request body is wrong, ", zap.Error(err))
		c.JSON(422, gin.H{"code": "REQUEST_BODY_IS_WRONG", "message": fmt.Sprintf("Request body is wrong. %s", err.Error())})
		return
	}

	c.JSON(200, gin.H{"code": "SUCCESS", "message": "Not Implemented"})
}

func (r *templateRoutes) postMethodTemplate(c *gin.Context) {

}
