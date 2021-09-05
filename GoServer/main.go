package main

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/controllers"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	models.ConnectDataBase()

	//Available endpoints
	r.GET("/listings", controllers.GetAllListings)
	r.POST("/listings", controllers.CreateListing)
	r.GET("/listings/:item_id", controllers.GetListingByItemID)
	r.PATCH("/listings/:item_id", controllers.UpdateListing)
	r.DELETE("/listings/:item_id")

	r.Run()
}
