package main

import (
	"bvdwalt/go-react/embedfs"
	"embed"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed _ui/build
var UI embed.FS

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	router.Use(static.Serve("/", embedfs.EmbedFolder(UI, "_ui/build", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "API")
		})
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
