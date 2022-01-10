package listings

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloWorld1(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	r.POST("v2/get_latest_listings", GetLatestListingsv2)

	params := url.Values{}
	params.Add("user_id", "5")
	params.Add("pagination", "1")
	paramString := params.Encode()

	req, _ := http.NewRequest("POST", "v2/get_latest_listings", strings.NewReader(paramString))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	// t.Run("get json data", func(t *testing.T) {
	// 	assert.Equal(t, 300, w.Code)
	// })
}
