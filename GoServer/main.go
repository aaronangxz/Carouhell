package main

import (
	"net/http"
	"os"

	"github.com/aaronangxz/TIC2601/auth"
	"github.com/aaronangxz/TIC2601/controllers/account"
	"github.com/aaronangxz/TIC2601/controllers/listings"
	"github.com/aaronangxz/TIC2601/controllers/wallet"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.LoadEnv() //loading env
	models.NewMySQL()
	models.NewAWSInstance()
	models.NewRedis()

	//Allow all CORS
	r.Use(cors.Default())
	r.Use(models.CORSMiddleware())

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
		c.JSON(http.StatusOK, gin.H{"Version": os.Getenv("CURRENT_VER"), "DockerImageID": os.Getenv("CURRENT_IMG_ID"), "LastUpdate": os.Getenv("LAST_UPDATE")})
	})

	//Test endpoint
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"RespMeta": "Hello World!"})
	})

	//Available endpoints
	// r.POST("/get_notifications_by_user_id", notifications.GetNotificationsByUserID)
	// r.POST("/create_mock_notifications", notifications.CreateMockNotifications)
	// r.POST("/upload_image", utils.UploadImage)

	//***************** Home Page *****************
	r.POST("/create_listing", auth.TokenAuthMiddleware(), listings.CreateListing)
	r.POST("/create_account", account.CreateAccount)
	r.POST("/authenticate_user", account.AuthenticateUser)

	//returns all, sorted by listing_ctime ASC
	r.GET("/get_all_listings", listings.GetAllListings)
	//returns all, sorted by listing_ctime DESC
	r.GET("/get_latest_listings", listings.GetLatestListings)
	//returns based on filters, sorted by listing_ctime DESC
	r.POST("/get_listings_using_filters", listings.GetListingsUsingFilters)

	//***************** Listing Page *****************
	//returns based on item_id
	r.POST("/get_single_listing_by_itemid", listings.GetListingByItemID)

	r.PATCH("/update_single_listing", auth.TokenAuthMiddleware(), listings.UpdateSingleListing)
	r.DELETE("/delete_single_listing", auth.TokenAuthMiddleware(), listings.DeleteListing)
	r.POST("/get_listing_reactions", listings.GetListingReactions)
	r.POST("/add_listing_likes", auth.TokenAuthMiddleware(), listings.AddListingLikes)
	r.POST("/add_listing_comments", auth.TokenAuthMiddleware(), listings.AddListingComments)
	r.POST("/purchase_single_item", auth.TokenAuthMiddleware(), listings.PurchaseSingleItem)

	//***************** Profile Page *****************
	//returns based on user_id, sorted by listing_ctime DESC
	r.POST("/get_user_listings", listings.GetUserListings)
	r.POST("/add_user_review", auth.TokenAuthMiddleware(), account.AddUserReview)

	//get_user_reviews
	r.POST("/get_user_details", account.GetUserDetails)

	//***************** Like Page *****************
	r.POST("/get_user_liked_listings", auth.TokenAuthMiddleware(), listings.GetUserLikedListings)

	//***************** Wallet Page *****************
	r.POST("/create_user_wallet", wallet.CreateUserWallet)
	r.POST("/get_user_wallet_details", auth.TokenAuthMiddleware(), wallet.GetUserWalletDetails)
	r.POST("/top_up_user_wallet", auth.TokenAuthMiddleware(), wallet.TopUpUserWallet)

	r.Run()
}
