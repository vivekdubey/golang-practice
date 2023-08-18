package password

import (
	"testing"
	"vivekdubey/golang-practice/internal/util"

	"github.com/stretchr/testify/require"
)

func countCharacters(cs []string, t string) int {
	count := 0
	for i := 0; i < len(t); i++ {
		if util.Contains(cs, string(t[i])) {
			count++
		}
	}
	return count
}

func TestRandomPassword(t *testing.T) {
	// check if addition of all minimums is more than total length
	_, err := RandomPassword(20, 3, 10, 10)
	require.Error(t, err)
	l := 20
	mc := 10
	mn := 3
	ms := 6
	v, err := RandomPassword(l, mn, mc, ms)
	require.NoError(t, err)
	cc := countCharacters(alphabetList(), v)
	cn := countCharacters(numberList(), v)
	cs := countCharacters(specialCharactersList(), v)
	require.GreaterOrEqual(t, cc, mc)
	require.GreaterOrEqual(t, cn, mn)
	require.GreaterOrEqual(t, cs, ms)
	require.Equal(t, l, len(v))
}
