package routers

import (
	"car_sales_mod/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {

	handler := handlers.NewHealthHandler()

	r.GET("/" , handler.Health)
	r.GET("/:id" , handler.Health_by_id)

}


