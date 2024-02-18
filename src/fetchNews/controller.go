package fetchNews

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsSharing/src/guardian"
	"newsSharing/src/nyTimes"
)

func NewsController(ginContext *gin.Context) {

	queryString, page := ginContext.Query("q"), ginContext.Query("page")

	fmt.Println(queryString)

	guardianChannel, nyTimesChannel := make(chan *guardian.NewsData, 1), make(chan *nyTimes.NewsData, 1)

	go guardian.GetNewsData(queryString, page, guardianChannel)
	go nyTimes.GetNewsData(queryString, page, nyTimesChannel)

	nyTimesData := <-nyTimesChannel
	guardianData := <-guardianChannel

	apiResponse := FormatNews(guardianData, nyTimesData)

	var (
		finalResult []byte
		err         error
	)

	if finalResult, err = json.Marshal(apiResponse); err != nil {
		// handle error
	}

	ginContext.Data(http.StatusOK, "application/json; charset=utf-8", finalResult)
}
