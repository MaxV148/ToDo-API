package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	userid := int64(12)
	token, err := GenerateToken(userid)
	fmt.Println(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

}
