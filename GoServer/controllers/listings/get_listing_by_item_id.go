package listings

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetListingByItemID(c *gin.Context) {
	var (
		resp     models.GetSingleListingResponse
		input    models.GetSingleListingRequest
		comments []models.ListingReactionsComments
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

	if input.GetItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be > 0.")})
		return
	}

	//also return deleted and sold items
	query := fmt.Sprintf("%v AND l.l_item_id = %v", utils.GetListingQueryWithCustomCondition(), input.GetItemID())
	log.Println(query)
	result := models.DB.Raw(query).Scan(&resp)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetListingByItemID - listing DB query: %v\n", err.Error())
		return
	}

	if resp.LItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
		log.Printf("item not found: item_id: %v", input.GetItemID())
		return
	}

	//retrieve comments
	queryComments := "SELECT a.user_name, l.comment, l.ctime FROM listing_reactions_tab l, acc_tab a " +
		"WHERE l.rt_user_id = a.a_user_id AND l.rt_item_id = " + fmt.Sprint(input.GetItemID()) + " AND l.reaction_type = 1 ORDER BY l.ctime ASC"
	log.Println(queryComments)
	resultComments := models.DB.Raw(queryComments).Scan(&comments)
	errComments := resultComments.Error

	if errComments != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errComments)})
		log.Printf("Error during GetListingByItemID - comments DB query: %v\n", errComments.Error())
		return
	}

	resp.Comments = comments
	resp.CommentCount = uint32(len(comments))

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v\n", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingByItemIDResult(resp), "Data": resp})
	log.Printf("Successful: GetListingByItemID. rows: %v\n", result.RowsAffected)
	log.Printf("Result: %s\n", data)
}
