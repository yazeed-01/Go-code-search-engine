package routes

import (
	"cse/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func UploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/search", Home)
	r.GET("/upload", UploadPage)

	r.POST("/upload", controllers.HandleFileUpload)
	r.POST("/search", controllers.Search)

	return r
}
