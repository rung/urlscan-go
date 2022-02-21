package urlscan_test

import (
	"encoding/json"
	"github.com/rung/urlscan-go/urlscan"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	jsonPath := os.Getenv("SCANURL_RESULT_JSON")
	if jsonPath == "" {
		t.Skip()
	}

	fd, err := os.Open(jsonPath)
	require.NoError(t, err)
	defer fd.Close()

	buf, err := ioutil.ReadAll(fd)
	require.NoError(t, err)

	var result urlscan.ScanResult
	err = json.Unmarshal(buf, &result)
	require.NoError(t, err)
}
