package listings

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aaronangxz/TIC2601/auth"
	"github.com/aaronangxz/TIC2601/models"
	"github.com/aaronangxz/TIC2601/utils"
	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

func GetListingByItemID(c *gin.Context) {
	var (
		resp          models.GetSingleListingResponse
		respWithLikes models.GetSingleListingLoggedInResponse
		input         models.GetSingleListingRequest
		comments      []models.ListingReactionsComments
		userId        uint64
		isLoggedIn    bool
		key           string
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ItemID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id cannot be empty.")})
			return
		}
		if !utils.ValidateUint(input.ItemID) {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be uint type.")})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewJSONErrorResponse(err)})
		return
	}

	if input.GetItemID() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewParamErrorsResponse("item_id must be > 0.")})
		return
	}

	//check token
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		log.Printf("Error during ExtractTokenMetadata: %v", err)
		isLoggedIn = false
	} else {
		//logged in
		userId, err = auth.FetchAuth(tokenAuth)
		if err != nil {
			log.Printf("Error during FetchAuth: %v, user is not logged in.\n", err)
		}
		log.Printf("user_id requesting is %v\n", userId)
		isLoggedIn = true
	}

	if !isLoggedIn {
		//not logged in

		//Redis key for not logged in
		key = fmt.Sprint("get_single_listing_by_itemid:", input.GetItemID())

		//check redis
		val, err := models.RedisClient.Get(models.Ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				log.Printf("No result of %v in Redis, reading from DB", key)
			} else {
				log.Printf("Error while reading from redis: %v", err.Error())
			}
		} else {
			redisResp := models.GetSingleListingResponse{}
			err := json.Unmarshal([]byte(val), &redisResp)
			if err != nil {
				log.Printf("Fail to unmarshal Redis value of key %v : %v, reading from DB", key, err)
			} else {
				c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingByItemIDResult(redisResp), "Data": redisResp})
				log.Printf("Successful: GetListingByItemID: %v - Cached", input.GetItemID())
				log.Printf("Result: %v\n", redisResp)
				return
			}
		}

		//Read from DB
		query := fmt.Sprintf("%v AND l.l_item_id = %v", utils.GetListingQueryWithCustomCondition(), input.GetItemID())
		log.Println(query)

		result := models.DB.Raw(query).Scan(&resp)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetListingByItemID DB query: %v\n", err.Error())
			return
		}

		//no result
		if resp.LItemID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
			log.Printf("item not found: item_id: %v", input.GetItemID())
			return
		}
	} else {
		//logged in
		//Redis key
		key = fmt.Sprintf("get_single_listing_by_itemid:%v_user:%v", input.GetItemID(), userId)

		//check redis
		val, err := models.RedisClient.Get(models.Ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				log.Printf("No result of %v in Redis, reading from DB", key)
			} else {
				log.Printf("Error while reading from redis: %v", err.Error())
			}
		} else {
			redisResp := models.GetSingleListingLoggedInResponse{}
			err := json.Unmarshal([]byte(val), &redisResp)
			if err != nil {
				log.Printf("Fail to unmarshal Redis value of key %v : %v, reading from DB", key, err)
			} else {
				c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingByItemIDLoggedInResult(redisResp), "Data": redisResp})
				log.Printf("Successful: GetListingByItemID: %v - Cached:user_id:%v", input.GetItemID(), userId)
				log.Printf("Result: %v\n", redisResp)
				return
			}
		}

		//Read from DB
		query := fmt.Sprintf("%v AND l.l_item_id = %v", utils.GetListingLoggedInQueryWithCustomCondition(userId), input.GetItemID())
		log.Println(query)
		result := models.DB.Raw(query).Scan(&respWithLikes)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(err)})
			log.Printf("Error during GetListingByItemID - listing DB query: %v\n", err.Error())
			return
		}

		if respWithLikes.LItemID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewNotFoundResponse()})
			log.Printf("item not found: item_id: %v", input.GetItemID())
			return
		}
	}

	//retrieve comments
	queryComments := "SELECT a.user_name, l.comment, l.ctime FROM listing_reactions_tab l, acc_tab a " +
		"WHERE l.rt_user_id = a.a_user_id AND l.rt_item_id = " + fmt.Sprint(input.GetItemID()) + " AND l.reaction_type = 1 ORDER BY l.ctime ASC"
	log.Println(queryComments)
	resultComments := models.DB.Raw(queryComments).Scan(&comments)
	errComments := resultComments.Error

	if errComments != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Respmeta": models.NewDBErrorResponse(errComments)})
		log.Printf("Error during GetListingByItemID - comments DB query: %v\n", errComments.Error())
		return
	}

	if !isLoggedIn {
		resp.Comments = comments
		resp.CommentCount = uint32(len(comments))

		data, err := json.Marshal(resp)
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

		c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingByItemIDResult(resp), "Data": resp})
		log.Printf("Successful: GetListingByItemID: %v - DB", input.GetItemID())
		log.Printf("Result: %s\n", data)
	} else {
		respWithLikes.Comments = comments
		respWithLikes.CommentCount = uint32(len(comments))

		data, err := json.Marshal(respWithLikes)
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

		c.JSON(http.StatusOK, gin.H{"Respmeta": utils.ValidateGetListingByItemIDLoggedInResult(respWithLikes), "Data": respWithLikes})
		log.Printf("Successful: GetListingByItemID: %v - DB:user_id:%v", input.GetItemID(), userId)
		log.Printf("Result: %s\n", data)
	}

}
