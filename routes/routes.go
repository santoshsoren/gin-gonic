package routes
import (
	//"log"
	"net/http"
	"github.com/gin-gonic/gin"
	controllers "github.com/santoshsoren/gin-gonic/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/allproduct", controllers.GetAllProducts)
	router.POST("/product", controllers.InsertProducts)
	router.GET("/product/:productId", controllers.GetSingleProduct)
	router.PUT("/product/:productId", controllers.UpdateProduct)
	router.DELETE("/product/:productId", controllers.DeleteProduct)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}