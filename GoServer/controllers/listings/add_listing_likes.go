package listings

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetLatestLikes(c *gin.Context, input models.AddListingLikeRequest) (uint32, error) {
	var count uint32

	//get current likes
	resultCount := models.DB.Table("listing_reactions_tab").Where("reaction_type = 0 AND rt_item_id = ?", input.GetItemID()).Count(&count)
	errCount := resultCount.Error

	if errCount != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errCount)})
		log.Printf("Error during AddListingLikes - get_current_likes DB query: %v\n", errCount.Error())
		errormsg := fmt.Sprintf("Error during AddListingLikes - get_current_likes DB query: %v\n", errCount.Error())
		return 0, errors.New(errormsg)
	}
	return count, nil
}

func AddListingLikes(c *gin.Context) {
	var (
		input        models.AddListingLikeRequest
		updatedLikes models.AddListingLikeResponse
		hold         models.ListingReactions
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
	reaction := models.ListingReactions{
		RTUserID:     input.UserID,
		RTItemID:     input.ItemID,
		ReactionType: utils.Uint32(constant.LISTING_REACTION_TYPE_LIKE),
		Comment:      nil,
		Ctime:        utils.Uint32(uint32(time.Now().Unix())),
	}

	//check if exists, if yes, delete record
	query := fmt.Sprintf("SELECT * FROM listing_reactions_tab WHERE rt_user_id = %v AND rt_item_id = %v AND reaction_type = %v", input.GetUserID(), input.GetItemID(), constant.LISTING_REACTION_TYPE_LIKE)
	find := models.DB.Raw(query).Scan(&hold)

	if find.Error != nil {
		if !errors.Is(find.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(find.Error)})
			log.Printf("Error during DB query: %v", find.Error.Error())
			return
		}
	}

	//delete record
	if find.RowsAffected > 0 {
		deleteQuery := fmt.Sprintf("DELETE FROM listing_reactions_tab WHERE rt_user_id = %v AND rt_item_id = %v AND reaction_type = %v", input.GetUserID(), input.GetItemID(), constant.LISTING_REACTION_TYPE_LIKE)
		if err := models.DB.Exec(deleteQuery).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during AddListingLikes DB - roll_back_likes query: %v\n", err.Error())
			return
		}
		log.Printf("already liked, rolled back record")

		//get latest likes
		count, err := GetLatestLikes(c, input)
		if err != nil {
			return
		}

		//return updated count
		updatedLikes.IsLiked = false
		updatedLikes.LikesCount = count
		log.Println("Successful: AddListingLikes.")
		c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse(fmt.Sprintf("Successfully decreased 1 like from listing %v", input.GetItemID())), "Data": updatedLikes})
		return
	}

	//insert like record
	result := models.DB.Table("listing_reactions_tab").Create(&reaction)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddListingLikes DB query: %v\n", err.Error())
		return
	}

	//get latest likes
	count, err := GetLatestLikes(c, input)
	if err != nil {
		return
	}

	updatedLikes.IsLiked = true
	updatedLikes.LikesCount = count

	log.Println("Successful: AddListingLikes.")
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse(fmt.Sprintf("Successfully added 1 like to listing %v", input.GetItemID())), "Data": updatedLikes})
}
