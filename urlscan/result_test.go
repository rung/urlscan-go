package urlscan

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetResult(t *testing.T) {
	uuid := "c623953a-74ed-408b-b3ba-22cb6eb2aaaa"
	_, err := GetResult(uuid)

	require.NoError(t, err)
}
