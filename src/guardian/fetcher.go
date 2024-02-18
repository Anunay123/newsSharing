package guardian

import (
	"encoding/json"
	"net/http"
	"newsSharing/src/network"
)

func GetNewsData(query, page string, newsDataChannel chan *NewsData) {
	queryMap := map[string]string{
		"q":       query,
		"api-key": "443a6aae-cb2c-4465-af1e-345cba365578",
		"page":    page,
	}

	responseBytes, err, _ := network.CallRestService(http.MethodGet, baseUrl, "Json", "Json",
		nil, queryMap, nil, nil, http.Client{})

	if err != nil {
		//error handling
		return
	}

	var newsData *NewsData

	if err := json.Unmarshal(responseBytes, &newsData); err != nil {
		//log unmarshal error
		return
	}

	newsDataChannel <- newsData
}
