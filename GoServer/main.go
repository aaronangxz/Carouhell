package main

import (
	"net/http"

	"github.com/aaronangxz/TIC2601/controllers/listings"
	"github.com/aaronangxz/TIC2601/controllers/notifications"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.LoadEnv() //loading env
	models.NewMySQL()

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
		c.JSON(http.StatusOK, gin.H{"Version": "1.6", "DockerImageID": "d85840c3a753", "LastUpdate": "03/10/2021 19:02"})
	})

	//Test endpoint
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"RespMeta": "Hello World!"})
	})

	//Available endpoints
	r.POST("/get_notifications_by_user_id", notifications.GetNotificationsByUserID)
	r.POST("/create_mock_notifications", notifications.CreateMockNotifications)

	r.GET("/get_all_listings", listings.GetAllListings)
	r.POST("/create_listing", listings.CreateListing)
	r.POST("/get_single_listing_by_itemid", listings.GetListingByItemID)
	r.PATCH("/update_single_listing", listings.UpdateSingleListing)
	r.DELETE("/delete_single_listing", listings.DeleteListing)
	r.POST("/get_user_listings", listings.GetUserListings)
	r.GET("/get_latest_listings", listings.GetLatestListings)
	r.POST("/get_listings_using_filters", listings.GetListingsUsingFilters)

	r.Run()
}
