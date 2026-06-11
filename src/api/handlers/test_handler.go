package handlers

import (
	"car_sales_mod/src/api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct{

	Firstname   string `json:"firstname" binding:"required,min=3,max=10"`
	Email 		string `json:"email" binding:"required,email,min=5,max=20"`
	Phone 	  	string `json:"phonenumber" binding:"required,numeric,iranphone,len=11"`

}

type testHeaderBinder struct {
	Userid		string
	Browser		string
}

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

func (h *TestHandler) Test_headerbineder1(c *gin.Context){

	user_id := c.GetHeader("user_id")

	c.JSON(http.StatusOK, gin.H{
		"result" : "Test_headerbineder1 !",
		"success" : true,
		"user_id" : user_id,
	})
}

func (h *TestHandler) Test_headerbineder2(c *gin.Context) {

	header := testHeaderBinder{}

	err := c.BindHeader(&header)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateBaseResponeWithError(nil,false,-1,err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseRespone("TestBinder2 called !",true , 0))
}

func (h *TestHandler) Test_querybinder1(c *gin.Context) {

	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "Test_querybinder1 !",
		"success":true,
		"id":id ,
		"name":name ,

	})
}

func (h *TestHandler) Test_querybinder2(c *gin.Context) {


	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "Test_querybinder1 !",
		"success":true,
		"ids":ids ,
		"name":name ,

	})
}

func (h *TestHandler) Test_bodybinder(c *gin.Context) {

	person := Person{}

	err := c.ShouldBindJSON(&person)

	if err != nil {
		c.AbortWithStatusJSON(403 , helper.GenerateBaseResponeWithValidationError(nil , false , -1, err))
	}

	c.JSON(http.StatusOK, helper.GenerateBaseRespone(
		[2]any{"Test_bodybinder !",person},
		true,
		0,
	))
	}


	//  if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error" : err.Error(),
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": "Test_bodybinder !",
	// 	"success":true,
	// 	"bodybinder": person,

	// })


func (h *TestHandler) Test_formbinder(c *gin.Context) {

	person := Person{}

	c.ShouldBind(&person)
	c.JSON(http.StatusOK, gin.H{
		"result": "Test_bodybinder !",
		"success":true,
		"formbinder": person,

	})
}


func (h *TestHandler) Test_filebinder(c *gin.Context) {

	file , _ := c.FormFile("file")
	err := c.SaveUploadedFile(file,"file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Test_filebinder !",
		"success":true,
		"file": file.Filename,

	})
}
