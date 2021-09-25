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
	r.Use(cors.Default())
	// r.LoadHTMLGlob("./home/*.html")
	// r.Static("/assets", "./home/assets")
	// r.Static("/images", "./home/images")
	// r.LoadHTMLGlob("./solution/*.html")
	// r.Static("/css", "./solution/css")
	// r.Static("/fonts", "./solution/fonts")
	// r.Static("/images", "./solution/images")
	// r.Static("/js", "./solution/js")
	// r.Static("/own-carousel", "./solution/own-carousel")
	// r.Static("/assets", "./solution/own-carousel/assets")

	r.LoadHTMLGlob("./elate/*.html")
	r.Static("/css", "./elate/css")
	r.Static("/fonts", "./elate/fonts")
	r.Static("/images", "./elate/images")
	r.Static("/js", "./elate/js")
	r.Static("/sass", "./elate/sass")
	r.Static("/syntax", "./elate/syntax")

	//Home endpoint
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
	})

	//Test endpoint
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"RespMeta": "Hello World!"})
	})

	models.ConnectDataBase()

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
