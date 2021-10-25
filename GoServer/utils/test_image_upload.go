package utils

import (
	"log"
	"net/http"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/gin-gonic/gin"
)

type UploadImageRequest struct {
	ItemID *uint32 `json:"item_id" binding:"required"`
	Base64 *string `json:"base64" binding:"required"`
}

func (r *UploadImageRequest) GetItemID() uint32 {
	if r != nil && r.ItemID != nil {
		return *r.ItemID
	}
	return 0
}

func (r *UploadImageRequest) GetBase64() string {
	if r != nil && r.Base64 != nil {
		return *r.Base64
	}
	return ""
}

type UploadImageResponse struct {
	URL string `json:"url"`
}

func UploadImage(c *gin.Context) {
	var (
		input UploadImageRequest
		resp  UploadImageResponse
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.Base64 == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("image cannot be empty.")})
			return
		}
	}

	url, err := UploadBase64Image(input.GetItemID(), input.GetBase64())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("Failed to upload image.")})
		log.Println(err)
		return
	}

	resp.URL = url

	log.Printf("Successful: UploadImage.")
	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully uploaded image"), "Data": resp})
}
