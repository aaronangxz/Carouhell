package account

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {
	var (
		input       models.GetUserDetailsRequest
		userDetails models.Account
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

	query := fmt.Sprintf("SELECT * FROM acc_tab WHERE user_id = %v", input.GetUserID())
	log.Println(query)

	result := models.DB.Raw(query).Scan(&userDetails)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserDetails DB query: %v\n", err.Error())
		return
	}

	log.Printf("Successful: GetUserDetails. rows: %v\n", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieve user details."), "Data": userDetails})
}
