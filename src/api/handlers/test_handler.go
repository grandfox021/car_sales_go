package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {

}


func NewTestHandler() *TestHandler{
	return &TestHandler{}
} 


func (h *TestHandler) TestHandler(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{
		"result" : "test handler called !",
		"success" : true,
	})
}


func (h *TestHandler) Test_user_handler_by_id(c *gin.Context){

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"result" : "test handler called !",
		"called_id" : id,
		"success" : true,
	})
}

func (h *TestHandler) Test_user_handler_post_by_username(c *gin.Context){

	username := c.Param("username")
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"result" : "test handler called !",
		"called_username" : username,
		"called_id" : id,
		"success" : true,
	})
}
