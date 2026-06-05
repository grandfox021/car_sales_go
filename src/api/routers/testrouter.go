package routers

import (
	"car_sales_mod/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func Testrouter(r *gin.RouterGroup){
	test_handler := handlers.NewTestHandler()
	r.GET("/" , test_handler.TestHandler)
	r.GET("/user/:id" , test_handler.Test_user_handler_by_id)
	r.POST("/user/:username/post/:id", test_handler.Test_user_handler_post_by_username)
	


}