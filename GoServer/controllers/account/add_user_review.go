package account

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/gin-gonic/gin"
)

func GetLatestRatings(c *gin.Context, input models.AddUserReviewRequest) (float32, error) {
	var (
		reviews models.AddUserReviewResponse
	)

	//get current likes
	query := fmt.Sprintf("SELECT ROUND((SUM(ratings)/ COUNT(ratings)) ,1) AS ratings FROM user_review_tab WHERE rv_seller_id = %v", input.GetSellerID())
	result := models.DB.Raw(query).Scan(&reviews)
	err := result.Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddUserReview - GetLatestRatings DB query: %v\n", err.Error())
		errormsg := fmt.Sprintf("Error during AddUserReview - GetLatestRatings DB query: %v\n", err.Error())
		return 0, errors.New(errormsg)
	}
	return reviews.Ratings, nil
}

func isExist(c *gin.Context, input models.AddUserReviewRequest) bool {
	var (
		count uint32
	)

	result := models.DB.Table("user_review_tab").
		Where("rv_user_id = ? AND rv_seller_id = ?",
			input.GetUserID(), input.GetSellerID()).Count(&count)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddUserReview - isExist DB query: %v\n", err.Error())
		return true
	}

	if count > 0 {
		return true
	}
	return false
}

func ValidateAddUserReviewRequest(c *gin.Context, input *models.AddUserReviewRequest) error {
	if err := c.ShouldBindJSON(&input); err != nil {
		if input.RVUserID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id cannot be empty.")})
			errormsg := "user_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.RVUserID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be uint type.")})
			errormsg := fmt.Sprintf("user_id must be uint type. input: %v", input.GetUserID())
			return errors.New(errormsg)
		}
		if input.RVSellerID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id cannot be empty.")})
			errormsg := "seller_id cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.RVSellerID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("seller_id must be uint type.")})
			errormsg := fmt.Sprintf("seller_id must be uint type. input: %v", input.GetUserID())
			return errors.New(errormsg)
		}
		if input.Ratings == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("ratings cannot be empty.")})
			errormsg := "ratings cannot be empty"
			return errors.New(errormsg)
		}
		if !utils.ValidateUint(input.Ratings) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("ratings must be uint type.")})
			errormsg := fmt.Sprintf("ratings must be uint type. input: %v", input.GetUserID())
			return errors.New(errormsg)
		}
		if input.ReviewText == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("review_text cannot be empty.")})
			errormsg := "review_text cannot be empty"
			return errors.New(errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		errormsg := fmt.Sprint("JSON error: &v", err.Error())
		return errors.New(errormsg)
	}
	return nil
}

func ValidateAddUserReviewInput(c *gin.Context, input *models.AddUserReviewRequest) error {
	if !utils.ValidateMaxStringLength(input.GetReviewText()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("review_text cannot exceed " + fmt.Sprint(models.MaxStringLength) + " chars.")})
		errormsg := fmt.Sprintf("review_text length cannot exceed %v. input :%v", models.MaxStringLength, len(input.GetReviewText()))
		return errors.New(errormsg)
	}
	//check if exists
	if input.Ratings != nil && !constant.CheckAccConstant(constant.USER_REVIEW_STARS, input.GetRatings()) {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("unknown rating.")})
		errormsg := fmt.Sprintf("unknown rating. input: %v", input.GetRatings())
		return errors.New(errormsg)
	}
	return nil
}

func AddUserReview(c *gin.Context) {
	var (
		input          models.AddUserReviewRequest
		updatedRatings models.AddUserReviewResponse
	)

	if err := ValidateAddUserReviewRequest(c, &input); err != nil {
		log.Printf("Error during ValidateAddUserReviewRequest: %v", err.Error())
		return
	}

	if err := ValidateAddUserReviewInput(c, &input); err != nil {
		log.Printf("Error during ValidateAddUserReviewInput: %v", err.Error())
		return
	}

	//check if reviewed before
	if isExist(c, input) {
		log.Printf("already reviewed before. user_id: %v, seller_id: %v", input.GetUserID(), input.GetSellerID())
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewUnknownErrorMessageResponse("Not allowed to review the same seller again.")})
		return
	}

	review := models.UserReview{
		RVUserID:   input.RVUserID,
		RVSellerID: input.RVSellerID,
		Ratings:    input.Ratings,
		ReviewText: input.ReviewText,
		Ctime:      utils.Int64(time.Now().Unix()),
	}

	if err := models.DB.Table("user_review_tab").Create(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during AddUserReview DB query: %v", err.Error())
		return
	}

	//get latest ratings
	count, err := GetLatestRatings(c, input)
	if err != nil {
		return
	}

	updatedRatings.Ratings = count

	if err := utils.InvalidateCache(utils.GetUserDetailsCacheKey, input.GetSellerID()); err != nil {
		log.Printf("Error during AddUserReview InvalidateCache: %v", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("Successfully added review."), "Data": updatedRatings})
	log.Println("Successful: AddUserReview.")
}
