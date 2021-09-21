package utils

import "github.com/aaronangxz/TIC2601/models"

func ValidateGetAllListingsResult(results []models.GetAllListingsResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessResponse()
}

func ValidateGetNotificationsByUserIDResult(results []models.GetNotificationsByUserIDResposne) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessResponse()
}
