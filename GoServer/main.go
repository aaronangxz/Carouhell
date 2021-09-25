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
	models.ConnectDataBase()

	//Allow all CORS
	r.Use(cors.Default())

	//Load HTML files for / endpoint
	r.LoadHTMLGlob("./elate/*.html")
	r.Static("/css", "./elate/css")
	r.Static("/fonts", "./elate/fonts")
	r.Static("/images", "./elate/images")
	r.Static("/js", "./elate/js")
	r.Static("/sass", "./elate/sass")
	r.Static("/syntax", "./elate/syntax")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
	})

	//Version endpoint
	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Version": "1.1", "DockerImageID": "9396ef44918a", "LastUpdate": "25/09/2021 02:34"})
	})

	//Test endpoint
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"RespMeta": "Hello World!"})
	})

	//Available endpoints
	r.GET("/get_notifications_by_user_id", controllers.GetNotificationsByUserID)
	r.POST("/create_mock_notifications", controllers.CreateMockNotifications)

	r.GET("/get_all_listings", controllers.GetAllListings)
	r.POST("/create_listing", controllers.CreateListing)
	r.GET("/get_single_listing_by_itemid", controllers.GetListingByItemID)
	r.PATCH("/update_single_listing", controllers.UpdateSingleListing)
	r.DELETE("/delete_single_listing", controllers.DeleteListing)
	r.GET("/get_user_listings", controllers.GetUserListings)

	r.Run()
}
