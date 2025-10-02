package example3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test global and nested error generation
// 测试全局和嵌套错误生成
func TestErrorGeneration(t *testing.T) {
	// Global error from errors.proto
	globalErr := ErrorUnknown("global error")
	require.True(t, IsUnknown(globalErr))
	t.Log(globalErr)

	// Nested error from user_service.proto
	nestedErr := ErrorGetUserResponseUserNotFound("user not found")
	require.True(t, IsGetUserResponseUserNotFound(nestedErr))
	t.Log(nestedErr)

	// Verify isolation
	require.False(t, IsGetUserResponseUserNotFound(globalErr))
	require.False(t, IsUnknown(nestedErr))
}

// Test same-name enum behavior
// 测试同名枚举的行为
func TestSameNameEnum(t *testing.T) {
	globalUnknown := ErrorUnknown("global")
	nestedUnknown := ErrorGetUserResponseUnknown("nested")

	// Same reason="UNKNOWN" makes them equivalent
	require.True(t, IsUnknown(globalUnknown))
	require.True(t, IsUnknown(nestedUnknown))
}
