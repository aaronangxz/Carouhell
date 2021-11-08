package account

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/auth"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {
	var (
		input        models.GetUserDetailsRequest
		accountResp  models.Account
		ratingsResp  models.UserRatings
		reviewsResp  []models.UserReviewWithNames
		listingsResp []models.GetUserListingsResponse
		userDetails  models.GetUserDetailsResponse
		mainstart    = time.Now().Unix()
		start        = int64(0)
		end          = int64(0)
		userId       uint64
		isLoggedIn   bool
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			log.Println("user_id cannot be empty.")
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		log.Printf("JSON error: %v\n", err.Error())
		return
	}

	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		//not logged in
		log.Printf("Error during ExtractTokenMetadata: %v", err)
		isLoggedIn = false
	} else {
		//logged in
		userId, err = auth.FetchAuth(tokenAuth)
		if err != nil {
			log.Printf("Error during FetchAuth: %v, user is not logged in.\n", err)
		}
		isLoggedIn = true
	}

	//retrieve acc info
	accQuery := fmt.Sprintf("SELECT * FROM acc_tab WHERE a_user_id = %v", input.GetUserID())
	log.Println(accQuery)
	start = time.Now().Unix()
	result := models.DB.Raw(accQuery).Scan(&accountResp)
	if err := result.Error; err != nil {
		end = time.Now().Unix()
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserDetails after %vs - retrieve acc info DB query: %v\n", end-start, err.Error())
		return
	}
	end = time.Now().Unix()
	log.Printf("Success: GetUserDetails after %vs - retrieve acc info DB query", end-start)

	//retrieve reviews
	reviewQuery := fmt.Sprintf("SELECT r.*, a.user_name AS user_name FROM user_review_tab r, acc_tab a WHERE r.rv_seller_id = %v AND r.rv_user_id= a.a_user_id ORDER BY ctime DESC", input.GetUserID())
	log.Println(reviewQuery)
	start = time.Now().Unix()
	result = models.DB.Raw(reviewQuery).Scan(&reviewsResp)
	if err := result.Error; err != nil {
		end = time.Now().Unix()
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserDetails after %vs- retrieve reviews DB query: %v\n", end-start, err.Error())
		return
	}
	end = time.Now().Unix()
	log.Printf("Success: GetUserDetails after %vs - retrieve reviews info DB query", end-start)

	if !isLoggedIn {
		//retrieve listings
		whereCondition := fmt.Sprintf(" AND l.l_seller_id = %v", input.GetUserID())
		orderCondition := " GROUP BY l.l_item_id ORDER BY listing_ctime DESC"
		listingQuery := utils.GetListingQueryWithCustomCondition() + whereCondition + orderCondition
		log.Println(listingQuery)
		start = time.Now().Unix()
		result = models.DB.Raw(listingQuery).Scan(&listingsResp)
		if err := result.Error; err != nil {
			end = time.Now().Unix()
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetUserDetails after %vs- retrieve listings DB query: %v\n", end-start, err.Error())
			return
		}
		end = time.Now().Unix()
		log.Printf("Success: GetUserDetails after %vs - retrieve listings info DB query", end-start)
	} else {
		whereCondition := fmt.Sprintf(" AND l.l_seller_id = %v", input.GetUserID())
		orderCondition := " GROUP BY l.l_item_id ORDER BY listing_ctime DESC"
		listingQuery := utils.GetListingLoggedInQueryWithCustomCondition(userId) + whereCondition + orderCondition
		log.Println(listingQuery)
		start = time.Now().Unix()
		result = models.DB.Raw(listingQuery).Scan(&listingsResp)
		if err := result.Error; err != nil {
			end = time.Now().Unix()
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetUserDetails after %vs- retrieve listings DB - logged in query: %v\n", end-start, err.Error())
			return
		}
		end = time.Now().Unix()
		log.Printf("Success: GetUserDetails after %vs - retrieve listings info DB - logged in query", end-start)
	}

	//get latest ratings
	ratingsQuery := fmt.Sprintf("SELECT ROUND((SUM(ratings)/ COUNT(ratings)) ,1) AS ratings FROM user_review_tab WHERE rv_seller_id = %v", input.GetUserID())
	start = time.Now().Unix()
	result = models.DB.Raw(ratingsQuery).Scan(&ratingsResp)
	if err := result.Error; err != nil {
		end = time.Now().Unix()
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserDetails after %vs- get latest ratings DB query: %v\n", end-start, err.Error())
		return
	}
	log.Printf("Success: GetUserDetails after %vs - get latest ratings DB query", end-start)

	//put into struct
	userDetails.AccountInfo = accountResp
	userDetails.Ratings = ratingsResp
	userDetails.ReviewCount = uint32(len(reviewsResp))
	userDetails.UserReviews = reviewsResp
	userDetails.ListingCount = uint32(len(listingsResp))
	userDetails.UserListings = listingsResp

	if listingsResp == nil {
		userDetails.UserListings = []models.GetUserListingsResponse{}
		userDetails.ListingCount = uint32(0)
	}

	mainend := time.Now().Unix()
	log.Printf("Successful: GetUserDetails after %vs. user_id: %v\n", mainend-mainstart, input.GetUserID())
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse(fmt.Sprintf("Successfully retrieved user details of %v", input.GetUserID())), "Data": userDetails})
}
