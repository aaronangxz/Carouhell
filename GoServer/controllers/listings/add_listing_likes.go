package listings

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func AddListingLikes(c *gin.Context) {
	var (
		input        models.AddListingLikeRequest
		count        uint32
		updatedLikes models.AddListingLikeResponse
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			return
		}
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	//write listing_reactions_tab
	// query := fmt.Sprintf("INSERT INTO listing_reactions_tab (user_id, item_id, reaction_type, ctime) VALUES (%v,%v,%v,%v)",
	// 	input.GetUserID(), input.GetItemID(), constant.LISTING_REACTION_TYPE_LIKE, time.Now().Unix())
	reaction := models.ListingReactions{
		UserID:       input.UserID,
		ItemID:       input.ItemID,
		ReactionType: utils.Uint32(constant.LISTING_REACTION_TYPE_LIKE),
		Comment:      nil,
		Ctime:        utils.Uint32(uint32(time.Now().Unix())),
	}
	//result := models.DB.Exec(query).Scan(&hold)
	result := models.DB.Table("listing_reactions_tab").Create(&reaction)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddListingLikes DB query: %v\n", err.Error())
		return
	}
	//get current likes
	resultCount := models.DB.Table("listing_reactions_tab").Where("reaction_type = 0 AND item_id = ?", input.GetItemID()).Count(&count)
	log.Println(resultCount)
	errCount := resultCount.Error

	if errCount != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errCount)})
		log.Printf("Error during AddListingLikes - get_current_likes DB query: %v\n", errCount.Error())
		return
	}

	updatedLikes.LikesCount = count

	log.Println("Successful: AddListingLikes.")
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse(fmt.Sprintf("Successfully added 1 like to listing %v", input.GetItemID())), "Data": updatedLikes})
}
