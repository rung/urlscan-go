package urlscan

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func GetDom(uuid string) (string, error) {
	resp, err := http.Get("https://urlscan.io/dom/" + uuid + "/")
	if err != nil {
		return "", errors.Errorf("Failed to get dom: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Errorf("Failed to get body: %s", err)
	}
	return string(body), nil
}
