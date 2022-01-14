package cart

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func ValidateDeleteItemFromUserCartInput(c *gin.Context, input *models.DeleteItemFromItemCartRequest) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.UserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("User ID cannot be empty.")})
			errormsg := "user_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateInt64(input.UserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("User ID must be int64 type.")})
			errormsg := fmt.Sprintf("user_id must be int64 type. input: %v", input.GetUserID())
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func DeleteItemFromUserCart(c *gin.Context) {
	var (
		input models.DeleteItemFromItemCartRequest
	)

	if err := ValidateDeleteItemFromUserCartInput(c, &input); err != nil {
		log.Printf("Error during ValidateDeleteItemFromUserCartInput: %v", err.Error())
		return
	}

	if err := models.DB.Exec("DELETE FROM user_cart_tab WHERE user_id = ? AND item_id = ?", input.GetUserID(), input.GetItemID()).Error; err != nil {
		log.Printf("Error during DeleteItemFromUserCart DB query: %v: %v", input.GetItemID(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successful: DeleteItemFromUserCart.")})
	log.Printf("Successfully deleted cart item: %v", input.GetItemID())
}
