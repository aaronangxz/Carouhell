package listings

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/constant"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

func GetUserLikedListings(c *gin.Context) {
	var (
		userLikedListings []models.GetUserLikedListingsResponse
		input             models.GetUserLikedListingsRequest
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

	if input.GetUserID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("user_id must be > 0.")})
		log.Println("user_id must be > 0.")
		return
	}

	//Redis key
	key := fmt.Sprint("get_user_liked_listings:", input.GetUserID())
	//check redis
	val, err := models.RedisClient.Get(models.Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("No result of %v in Redis, reading from DB", key)
		} else {
			log.Printf("Error while reading from redis: %v", err.Error())
		}
	} else {
		redisResp := []models.GetUserLikedListingsResponse{}
		err := json.Unmarshal([]byte(val), &redisResp)
		if err != nil {
			log.Printf("Fail to unmarshal Redis value of key %v : %v, reading from DB", key, err)
		}
		c.JSON(http.StatusOK, gin.H{"Respmeta": models.NewSuccessMessageResponse("GetUserLikedListings success."), "Data": redisResp})
		log.Printf("Successful: GetUserLikedListings - Cached: %v ", key)
		return
	}

	//also return deleted and sold items
	query := fmt.Sprintf(utils.GetListingLoggedInQueryWithCustomCondition(uint64(input.GetUserID()))+" AND l.l_item_id IN"+
		" (SELECT r.rt_item_id FROM listing_reactions_tab r"+
		" WHERE r.rt_user_id = %v AND r.reaction_type = %v)"+
		" GROUP BY l.l_item_id ORDER BY listing_reactions_tab.ctime DESC", input.GetUserID(), constant.LISTING_REACTION_TYPE_LIKE)
	log.Println(query)

	result := models.DB.Raw(query).Scan(&userLikedListings)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
		log.Printf("Error during GetUserLikedListings DB query: %v\n", err.Error())
		return
	}

	data, err := json.Marshal(userLikedListings)
	if err != nil {
		log.Printf("Failed to marshal JSON results: %v\n", err.Error())
	}

	//Retrieved new data, set to Redis
	expiry := 60 * time.Second
	if err := models.RedisClient.Set(models.Ctx, key, data, expiry).Err(); err != nil {
		log.Printf("Error while writing to redis: %v", err.Error())
	} else {
		log.Printf("Written %v to redis", key)
	}

	c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetUserLikedListingsResult(userLikedListings), "Data": userLikedListings})
	log.Printf("Successful: GetUserLikedListings - DB. rows: %v\n", result.RowsAffected)
	log.Printf("Result: %s\n", data)
}
