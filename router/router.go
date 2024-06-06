package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações Default do gin
	r := gin.Default()
	r.GET("/vai", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Utilizando Router!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
