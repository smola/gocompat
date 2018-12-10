package compat

import (
	"testing"

	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/require"
)

func TestAPIJSON(t *testing.T) {
	require := require.New(t)
	expected, err := ReachableFromPackages(".")
	require.NoError(err)
	require.True(len(expected.Packages) == 1)
	require.True(len(expected.Packages[0].Objects) > 0)

	bytes, err := easyjson.Marshal(expected)
	require.NoError(err)

	actual := NewAPI()
	err = easyjson.Unmarshal(bytes, actual)
	require.NoError(err)
	require.Equal(expected, actual)
}
