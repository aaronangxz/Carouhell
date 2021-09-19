package utils

import "github.com/aaronangxz/TIC2601/models"

func ValidateBatchListingResult(listings []models.Listing) models.ResponseMeta {
	if len(listings) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessResponse()
}
