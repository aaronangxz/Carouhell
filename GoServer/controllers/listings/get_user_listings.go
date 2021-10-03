package listings

import (
	"fmt"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetUserListings(c *gin.Context) {
	var (
		userListings   []models.Listing
		input          models.GetUserListingsRequest
		extraCondition = ""
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetUserID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be > 0.")})
		return
	}

	if input.Limit == nil {
		input.Limit = models.SetDefaultNotificationResponseLimit()
	}

	if utils.ValidateLimitMax(input.GetLimit(), models.MaxListingsResponseSize) {
		c.JSON(http.StatusBadRequest, gin.H{"RespMeta": models.NewParamErrorsResponse("limit exceeds max listing response size.")})
		return
	}

	switch input.GetLimit() {
	case 0:
		extraCondition = ""
	default:
		extraCondition = " LIMIT " + fmt.Sprint(input.GetLimit())
	}

	query := "SELECT * FROM listing_tab WHERE seller_id = ? ORDER BY listing_ctime DESC" + extraCondition

	if err := models.DB.Raw(query, input.GetUserID()).
		Scan(&userListings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	if len(userListings) == 0 {
		c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewNotFoundResponse(), "Data": userListings})
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessResponse(), "Data": userListings})
}
