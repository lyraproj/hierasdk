package hiera

import (
	"testing"

	"github.com/lyraproj/dgo/test/require"
	"github.com/lyraproj/dgo/vf"
)

func TestProviderContext_StringOption(t *testing.T) {
	c := ProviderContextFromMap(vf.Map(`s`, `a`, `i`, 2))
	s, ok := c.StringOption(`s`)
	require.True(t, ok)
	require.Equal(t, `a`, s)
	_, ok = c.StringOption(`i`)
	require.False(t, ok)
}

func TestProviderContext_IntOption(t *testing.T) {
	c := ProviderContextFromMap(vf.Map(`s`, `a`, `i`, 2))
	i, ok := c.IntOption(`i`)
	require.True(t, ok)
	require.Equal(t, 2, i)
	_, ok = c.IntOption(`s`)
	require.False(t, ok)
}

func TestProviderContext_FloatOption(t *testing.T) {
	c := ProviderContextFromMap(vf.Map(`s`, `a`, `f`, 2.0))
	f, ok := c.FloatOption(`f`)
	require.True(t, ok)
	require.Equal(t, 2.0, f)
	_, ok = c.FloatOption(`s`)
	require.False(t, ok)
}

func TestProviderContext_BoolOption(t *testing.T) {
	c := ProviderContextFromMap(vf.Map(`s`, `a`, `b`, false))
	b, ok := c.BoolOption(`b`)
	require.True(t, ok)
	require.Equal(t, false, b)
	_, ok = c.BoolOption(`s`)
	require.False(t, ok)
}

func TestProviderContext_OptionsMap(t *testing.T) {
	c := ProviderContextFromMap(nil)
	require.Equal(t, vf.Map(), c.OptionsMap())
}
