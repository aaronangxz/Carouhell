package account

import (
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

	query := "SELECT * FROM acc_tab WHERE user_id = ?"
	log.Println(query)

	result := models.DB.Raw(query, input.GetUserID).Scan(&userDetails)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserDetails DB query: %v\n", err.Error())
		return
	}

	log.Printf("Successful: GetAllListings. rows: %v\n", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully retrieve user details."), "Data": userDetails})
}
