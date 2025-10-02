package example2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNestedEnums(t *testing.T) {
	// Test nested enum in message
	err := ErrorGetInfoRespUserNotFound("user not found")
	require.True(t, IsGetInfoRespUserNotFound(err))
	t.Log(err)
}

func TestTopLevelEnums(t *testing.T) {
	// Test top-level enum
	err := ErrorUnknown("unknown error")
	require.True(t, IsUnknown(err))
	t.Log(err)

	err = ErrorServiceUnavailable("service unavailable")
	require.True(t, IsServiceUnavailable(err))
	t.Log(err)
}

func TestEnumIsolation(t *testing.T) {
	// Test isolation between enums
	topErr := ErrorUnknown("top level error")
	nestedErr := ErrorGetInfoRespUserNotFound("nested error")

	require.False(t, IsGetInfoRespUserNotFound(topErr))
	require.False(t, IsUnknown(nestedErr))
}
