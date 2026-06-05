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
	r.POST("/user/headerbind1",test_handler.Test_headerbineder1)
	r.POST("/user/headerbind2",test_handler.Test_headerbineder2)

	r.POST("/user/querybinder1",test_handler.Test_querybinder1)
	r.POST("/user/querybinder2",test_handler.Test_querybinder2)

	r.POST("/user/bodybinder",test_handler.Test_bodybinder)

	r.POST("/user/formbinder",test_handler.Test_formbinder)

	r.POST("/user/filebinder",test_handler.Test_filebinder)

}