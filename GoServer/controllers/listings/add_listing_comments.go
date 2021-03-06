package listings

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func isSpam(c *gin.Context, input models.AddListingCommentsRequest) bool {
	var (
		count uint32
	)

	//if antispam is off
	if os.Getenv("CONFIG_ANTISPAM") == "FALSE" {
		return false
	}

	result := models.DB.Table("listing_reactions_tab").
		Where(" reaction_type = ? AND rt_user_id = ? AND rt_item_id = ? AND ? - ctime <= 60",
			constant.LISTING_REACTION_TYPE_COMMENT, input.GetUserID(), input.GetItemID(), time.Now().Unix()).Count(&count)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddListingComments - isSpam DB query: %v\n", err.Error())
		return true
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("ANTISPAM toggle is on. Not allowed to comment again within 60 secs.")})
		return true
	}
	return false
}

func AddListingComments(c *gin.Context) {
	var (
		input models.AddListingCommentsRequest
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			log.Println("item_id cannot be empty.")
			return
		}
		if !utils.ValidateUint(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			log.Println("item_id must be uint type.")
			return
		}
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			log.Println("user_id cannot be empty.")
			return
		}
		if !utils.ValidateUint(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be uint type.")})
			log.Println("user_id must be uint type.")
			return
		}
		if input.Comment == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("comment cannot be empty.")})
			log.Println("comment canot be empty.")
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if !utils.ValidateMaxStringLength(input.GetComment()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("comment cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		log.Printf("comment length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetComment()))
		return
	}

	if input.GetComment() == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("comment canot be empty.")})
		log.Println("comment canot be empty.")
		return
	}

	//write listing_reactions_tab
	reaction := models.ListingReactions{
		RTUserID:     input.UserID,
		RTItemID:     input.ItemID,
		ReactionType: utils.Uint32(constant.LISTING_REACTION_TYPE_COMMENT),
		Comment:      input.Comment,
		Ctime:        utils.Uint32(uint32(time.Now().Unix())),
	}

	//spam prevention, not allowed to post again within 60 secs
	if isSpam(c, input) {
		log.Printf("comment spam detected. user_id: %v, item_id: %v", input.GetUserID(), input.GetItemID())
		return
	}

	//insert comment record
	result := models.DB.Table("listing_reactions_tab").Create(&reaction)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddListingComments DB query: %v\n", err.Error())
		return
	}

	//invalidate redis
	if err := utils.InvalidateCache(utils.GetSingleListingByUserIDCacheKey, input.GetItemID()); err != nil {
		log.Printf("Error during InvalidateCache: %v", err.Error())
	}

	log.Println("Successful: AddListingComments.")
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse(fmt.Sprintf("Successfully added comment to listing %v", input.GetItemID()))})
}
