package nyTimes

import (
	"encoding/json"
	"net/http"
	"newsSharing/src/network"
)

func GetNewsData(query, page string, newsDataChannel chan *NewsData) {
	queryMap := map[string]string{
		"q":       query,
		"api-key": "joSrRFx4DPYagMDdofeDeHQdUAvhGjXl",
		"page":    page,
	}

	responseBytes, err, _ := network.CallRestService(http.MethodGet, baseUrl, "Json", "Json",
		nil, queryMap, nil, nil, http.Client{})

	if err != nil {
		return
	}

	var newsData *NewsData

	if err := json.Unmarshal(responseBytes, &newsData); err != nil {
		//log unmarshal error
		return
	}
	newsDataChannel <- newsData
}
