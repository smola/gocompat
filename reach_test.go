package compat

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReachableFromPackages_Packages(t *testing.T)  {
	require := require.New(t)
	api, err := reachableFromPackages(true, ".")
	require.NoError(err)
	require.NotNil(api)
	require.GreaterOrEqual(len(api.Packages), 2)
}

func TestReachableFromPackages_Objects1(t *testing.T)  {
	require := require.New(t)
	api, err := ReachableFromPackages(".")
	require.NoError(err)
	require.NotNil(api)
	require.Equal(len(api.Packages), 1)
	require.GreaterOrEqual(len(api.Reachable), 10)
}

func TestReachableFromPackages_Objects2(t *testing.T)  {
	require := require.New(t)
	api, err := ReachableFromPackages("./...")
	require.NoError(err)
	require.NotNil(api)
	require.Equal(len(api.Packages), 2)
	require.GreaterOrEqual(len(api.Reachable), 10)
}
