package urlscan

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type Result struct {
	Verdicts struct {
		Overall struct {
			Score       int      `json:"score"`
			Categories  []string `json:"categories"`
			Brands      []string `json:"brands"`
			Tags        []string `json:"tags"`
			Malicious   bool     `json:"malicious"`
			HasVerdicts bool     `json:"hasVerdicts"`
		} `json:"overall"`
	} `json:"verdicts"`
}

func GetResult(uuid string) (*Result, error) {
	resp, err := http.Get("https://urlscan.io/api/v1/result/" + uuid + "/")
	if err != nil {
		return nil, errors.Errorf("Failed to get result: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Failed to get result: %s", err)
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Errorf("Failed to unmarshal result: %s", err)
	}

	return &result, nil
}
