package listings

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetListingReactions(c *gin.Context) {
	var (
		input     models.GetListingReactionsRequest
		reactions models.GetListingReactionsResponse
		count     uint32
		comments  []models.ListingReactionsComments
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
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	//total likes
	resultLikes := models.DB.Table("listing_reactions_tab").Where("reaction_type = 0 AND rt_item_id = ?", input.GetItemID()).Count(&count)
	errLikes := resultLikes.Error

	if errLikes != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errLikes)})
		log.Printf("Error during GetListingReactions - likes_count DB query: %v\n", errLikes.Error())
		return
	}

	//retrieve comments
	queryComments := "SELECT a.user_name, l.rt_user_id AS user_id, l.comment, l.ctime FROM listing_reactions_tab l, acc_tab a " +
		"WHERE l.rt_user_id = a.a_user_id AND l.rt_item_id = " + fmt.Sprint(input.GetItemID()) + " AND l.reaction_type = 1 ORDER BY l.ctime ASC"
	log.Println(queryComments)
	resultComments := models.DB.Raw(queryComments).Scan(&comments)
	errComments := resultComments.Error

	if errComments != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errComments)})
		log.Printf("Error during GetListingReactions - comment DB query: %v\n", errComments.Error())
		return
	}

	reactions.LikesCount = count
	reactions.Comments = comments

	log.Println("Successful: GetListingReactions.")
	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingReactionsResult(reactions.Comments), "Data": reactions})
}
