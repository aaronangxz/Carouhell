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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
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
