package utils

import (
	"fmt"

	"github.com/aaronangxz/TIC2601/models"
)

func ValidateGetAllListingsResult(results []models.GetAllListingsResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse(fmt.Sprintf("GetAllListings success. results: %v", len(results)))
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
	return models.NewSuccessMessageResponse(fmt.Sprintf("GetLatestListings success. results: %v", len(results)))
}

func ValidateGetListingsUsingFiltersResult(results []models.GetListingsUsingFiltersResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse(fmt.Sprintf("GetListingsUsingFilters success. results: %v", len(results)))
}

func ValidateGetUserListingsResult(results []models.GetUserListingsResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse(fmt.Sprintf("GetUserListings success. results: %v", len(results)))
}
