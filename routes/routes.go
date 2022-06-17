package routes

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signup", controllers.Signup())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", conroller.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearhProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByOrder())
}
