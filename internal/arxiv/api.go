package arxiv

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"newsly-ingestion/internal/utils"
)

type API struct {
	BaseUrl string
}

func InitArxivAPI(url string) *API {
	return &API{BaseUrl: url}
}

func (api *API) makeGetRequest(queryParams map[string]string) []byte {
	request, err := http.NewRequest("GET", api.BaseUrl, nil)

	q := request.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}

	request.URL.RawQuery = q.Encode()

	response, err := http.DefaultClient.Do(request)
	utils.PanicOnError(err, fmt.Sprintf("error fetching %s", api.BaseUrl))

	body, err := ioutil.ReadAll(response.Body)
	utils.PanicOnError(err, fmt.Sprintf("error reading data at %s", api.BaseUrl))
	return body
}

func (api *API) GetArticlesForTerm(term string) Feed {
	params := make(map[string]string)
	params["search_query"] = term
	params["sortBy"] = "submittedDate"
	params["sortOrder"] = "descending"

	body := api.makeGetRequest(params)

	var FeedResponse Feed

	err := xml.Unmarshal(body, &FeedResponse)
	utils.PanicOnError(err, "unable to unmarshall XML to Feed type")
	return FeedResponse
}
