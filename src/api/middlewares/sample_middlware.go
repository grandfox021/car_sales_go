package middlewares

import "github.com/gin-gonic/gin"

func AuthSampleMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context){
		key := ctx.GetHeader("x-api-key")
		if key == "1"{
			ctx.Next()
		}else{
		ctx.AbortWithStatusJSON(401 , gin.H{
			"result" : "api_key is required or not correct !",
		})

		return 
	}
  }
}