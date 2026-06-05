package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{

}

func NewHealthHandler() *HealthHandler {

	return &HealthHandler{}

}


func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(200,"working!!")
}

func (h *HealthHandler) Health_by_id(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(200,fmt.Sprintf("working!! id: %s" , id))
}