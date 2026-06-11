package api

import (
	"car_sales_mod/src/api/middlewares"
	"car_sales_mod/src/api/routers"
	"car_sales_mod/src/api/validations"
	"car_sales_mod/src/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


func Init_server(){

	cfg := config.GetConfig()

	r := gin.New()
	val , ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("iranphone" , validations.IranianPhonenumberValidator,true)
	}

	r.Use(gin.Logger(),gin.Recovery(),middlewares.AuthSampleMiddleware())

	api := r.Group("/api")
	v1 := api.Group("/v1")

	health := v1.Group("/health")
	test := v1.Group("/test")
	routers.Health(health)
	routers.Testrouter(test)
	
	// r.Run("localhost:5005")
	r.Run(fmt.Sprintf(":%s" , cfg.Server.Port))
	
}