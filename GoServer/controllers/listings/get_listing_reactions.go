package listings

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func GetListingReactions(c *gin.Context) {
	var (
		input     models.GetListingReactionsRequest
		reactions models.GetListingReactionsResponse
	)

	//total likes
	queryLikes := "SELECT COUNT(*) FROM listing_reactions_tab WHERE reaction_type = 0 AND item_id = " + fmt.Sprint(input.GetItemID())
	log.Println(queryLikes)
	resultLikes := models.DB.Raw(queryLikes).Scan(&reactions.LikesCount)
	errLikes := resultLikes.Error

	if errLikes != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errLikes)})
		log.Printf("Error during GetListingReactions - likes_count DB query: %v\n", errLikes.Error())
		return
	}

	//retrieve comments
	queryComments := "SELECT a.user_name, l.comment, l.ctime FROM listing_reactions_tab l, acc_tab a " +
		"WHERE l.user_id = a.user_id AND item_id = " + fmt.Sprint(input.GetItemID()) + " AND reaction_type = 1 ORDER BY ctime ASC"
	resultComments := models.DB.Raw(queryComments).Scan(&reactions.Comments)
	errComments := resultComments.Error

	if errLikes != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errComments)})
		log.Printf("Error during GetListingReactions - likes_count DB query: %v\n", errComments.Error())
		return
	}

	log.Println("Successful: GetListingReactions.")
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieve listing reactions."), "Data": reactions})
}
