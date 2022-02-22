package urlscan

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
)

// SearchArguments is input data structure of Search()
type SearchArguments struct {
	// Optional. urlscan.io search query.
	// See Help & Example of https://urlscan.io/search/ for more detail
	Query *string `json:"query"`
	// Optional. Page size
	Size *uint64 `json:"size"`
	// Optional.
	SearchAfter *string `json:"search_after"`
}

// SearchResult represents a single search result from the API
type SearchResult struct {
	IndexedAt string `json:"indexedAt"`
	ID        string `json:"_id"`
	Page      struct {
		Asn     string `json:"asn"`
		Asnname string `json:"asnname"`
		City    string `json:"city"`
		Country string `json:"country"`
		Domain  string `json:"domain"`
		IP      string `json:"ip"`
		Ptr     string `json:"ptr"`
		Server  string `json:"server"`
		URL     string `json:"url"`
	} `json:"page"`
	Result     string `json:"result"`
	Screenshot string `json:"screenshot"`
	Stats      struct {
		ConsoleMsgs       int64 `json:"consoleMsgs"`
		DataLength        int64 `json:"dataLength"`
		EncodedDataLength int64 `json:"encodedDataLength"`
		Requests          int64 `json:"requests"`
		UniqIPs           int64 `json:"uniqIPs"`
	} `json:"stats"`
	Task struct {
		Method     string `json:"method"`
		UUID       string `json:"uuid"`
		Source     string `json:"source"`
		Time       string `json:"time"`
		URL        string `json:"url"`
		Visibility string `json:"visibility"`
	} `json:"task"`
	RawSort       []interface{} `json:"sort"`
	UniqCountries int64         `json:"uniq_countries"`
}

// SearchResponse is returned by Search() and including existing scan results.
type SearchResponse struct {
	Results []SearchResult `json:"results"`
	Total   int64          `json:"total"`
}

// Search sends query to search existing scan results with query
func (x *Client) Search(args SearchArguments) (SearchResponse, error) {
	var result SearchResponse
	values := make(url.Values)

	if args.Query != nil {
		values.Add("q", *args.Query)
	}
	if args.Size != nil {
		values.Add("size", fmt.Sprintf("%d", *args.Size))
	}
	if args.SearchAfter != nil {
		values.Add("search_after", *args.SearchAfter)
	}

	code, err := x.get("search", values, &result)
	if err != nil {
		return result, err
	}
	if code != 200 {
		return result, errors.Errorf("Unexpected status code: %d", code)
	}

	return result, err
}

func NormalizeSort(rawSort []interface{}) (string, error) {
	number, ok := rawSort[0].(float64)
	if !ok {
		return "", errors.Errorf("Failed to normalize sort")
	}
	uuid, ok := rawSort[1].(string)
	if !ok {
		return "", errors.Errorf("Failed to normalize sort")
	}

	return fmt.Sprintf("%d,%s", uint64(number), uuid), nil
}
