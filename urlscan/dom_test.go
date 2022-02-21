package urlscan

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetDom(t *testing.T) {
	uuid := "fe0390af-48f6-47b0-b720-d4bab9977314"
	dom, err := GetDom(uuid)
	fmt.Println(dom)

	require.NoError(t, err)
}