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

func ValidateGetLatestListingsResult(results []models.GetLatestListingsResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse(fmt.Sprintf("GetLatestListings success. results: %v", len(results)))
}

func ValidateGetLatestListingsLoggedInResult(results []models.GetLatestListingsLoggedInResponse) models.ResponseMeta {
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

func ValidateGetListingsUsingFiltersLoggedInResult(results []models.GetListingsUsingFiltersLoggedInResponse) models.ResponseMeta {
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

func ValidateGetListingReactionsResult(results []models.ListingReactionsComments) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse(fmt.Sprintf("GetListingReactions success. results: %v", len(results)))
}

func ValidateGetUserLikedListingsResult(results []models.GetUserLikedListingsResponse) models.ResponseMeta {
	if len(results) == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse(fmt.Sprintf("GetUserLikedListings success. results: %v", len(results)))
}

func ValidateGetListingByItemIDResult(results models.GetSingleListingResponse) models.ResponseMeta {
	if results.LItemID == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse("GetListingByItemID success.")
}

func ValidateGetListingByItemIDLoggedInResult(results models.GetSingleListingLoggedInResponse) models.ResponseMeta {
	if results.LItemID == 0 {
		return models.NewNotFoundResponse()
	}
	return models.NewSuccessMessageResponse("GetListingByItemID success.")
}
