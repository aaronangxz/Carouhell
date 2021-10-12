package utils

import "github.com/aaronangxz/TIC2601/models"

func ValidateGetAllListingsResult(results []models.GetAllListingsResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessResponse()
}

func ValidateGetNotificationsByUserIDResult(results []models.GetNotificationsByUserIDResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessResponse()
}

func ValidateGetLatestListingsResult(results []models.GetLatestListingsResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessResponse()
}

func ValidateGetListingsUsingFiltersResult(results []models.Listing) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessResponse()
}
