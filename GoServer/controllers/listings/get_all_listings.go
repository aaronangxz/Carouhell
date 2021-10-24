package listings

import (
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"

	"github.com/gin-gonic/gin"
)

func GetAllListings(c *gin.Context) {
	var (
		listings []models.GetAllListingsResponse
	)

	query := "SELECT l.item_id, l.item_name, l.item_price, l.item_quantity," +
		" l.item_purchased_quantity, l.item_description, l.item_shipping_info," +
		" l.item_payment_info,l.item_location, l.item_status, l.item_category," +
		" l.item_image, l.seller_id, a.user_name AS seller_name, l.listing_ctime,l.listing_mtime, l.listing_likes" +
		" FROM listing_tab l, acc_tab a WHERE l.seller_id = a.user_id GROUP BY l.item_id"
	log.Println(query)

	result := models.DB.Raw(query).Scan(&listings)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetAllListings DB query: %v\n", err.Error())
		return
	}

	log.Printf("Successful: GetAllListings. rows: %v\n", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetAllListingsResult(listings), "Data": listings})
}
