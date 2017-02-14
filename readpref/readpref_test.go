package readpref_test

import (
	"testing"
	"time"

	. "github.com/10gen/mongo-go-driver/readpref"
	"github.com/10gen/mongo-go-driver/server"
	"github.com/stretchr/testify/require"
)

func TestPrimary(t *testing.T) {
	require := require.New(t)
	subject := Primary()

	require.Equal(PrimaryMode, subject.Mode())
	_, set := subject.MaxStaleness()
	require.False(set)
	require.Empty(subject.TagSets())
}

func TestPrimaryPreferred(t *testing.T) {
	require := require.New(t)
	subject := PrimaryPreferred()

	require.Equal(PrimaryPreferredMode, subject.Mode())
	_, set := subject.MaxStaleness()
	require.False(set)
	require.Empty(subject.TagSets())
}

func TestPrimaryPreferred_with_options(t *testing.T) {
	require := require.New(t)
	subject := PrimaryPreferred(
		WithMaxStaleness(time.Duration(10)),
		WithTags("a", "1"),
	)

	require.Equal(PrimaryPreferredMode, subject.Mode())
	ms, set := subject.MaxStaleness()
	require.True(set)
	require.Equal(time.Duration(10), ms)
	require.Equal([]server.TagSet{server.NewTagSet("a", "1")}, subject.TagSets())
}

func TestSecondaryPreferred(t *testing.T) {
	require := require.New(t)
	subject := SecondaryPreferred()

	require.Equal(SecondaryPreferredMode, subject.Mode())
	_, set := subject.MaxStaleness()
	require.False(set)
	require.Empty(subject.TagSets())
}

func TestSecondaryPreferred_with_options(t *testing.T) {
	require := require.New(t)
	subject := SecondaryPreferred(
		WithMaxStaleness(time.Duration(10)),
		WithTags("a", "1"),
	)

	require.Equal(SecondaryPreferredMode, subject.Mode())
	ms, set := subject.MaxStaleness()
	require.True(set)
	require.Equal(time.Duration(10), ms)
	require.Equal([]server.TagSet{server.NewTagSet("a", "1")}, subject.TagSets())
}

func TestSecondary(t *testing.T) {
	require := require.New(t)
	subject := Secondary()

	require.Equal(SecondaryMode, subject.Mode())
	_, set := subject.MaxStaleness()
	require.False(set)
	require.Empty(subject.TagSets())
}

func TestSecondary_with_options(t *testing.T) {
	require := require.New(t)
	subject := Secondary(
		WithMaxStaleness(time.Duration(10)),
		WithTags("a", "1"),
	)

	require.Equal(SecondaryMode, subject.Mode())
	ms, set := subject.MaxStaleness()
	require.True(set)
	require.Equal(time.Duration(10), ms)
	require.Equal([]server.TagSet{server.NewTagSet("a", "1")}, subject.TagSets())
}

func TestNearest(t *testing.T) {
	require := require.New(t)
	subject := Nearest()

	require.Equal(NearestMode, subject.Mode())
	_, set := subject.MaxStaleness()
	require.False(set)
	require.Empty(subject.TagSets())
}

func TestNearest_with_options(t *testing.T) {
	require := require.New(t)
	subject := Nearest(
		WithMaxStaleness(time.Duration(10)),
		WithTags("a", "1"),
	)

	require.Equal(NearestMode, subject.Mode())
	ms, set := subject.MaxStaleness()
	require.True(set)
	require.Equal(time.Duration(10), ms)
	require.Equal([]server.TagSet{server.NewTagSet("a", "1")}, subject.TagSets())
}