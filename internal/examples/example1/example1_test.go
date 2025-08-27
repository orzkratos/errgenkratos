package example1_test

import (
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errgenkratos"
	"github.com/orzkratos/errgenkratos/internal/examples/example1"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	errgenkratos.SetMetadataFieldName("reason_numeric")
	m.Run()
}

func TestErrorUnknown(t *testing.T) {
	erk := example1.ErrorUnknown("test unknown error: %s", "details")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, example1.ErrorReason_UNKNOWN.String(), erk.Reason)
	require.Equal(t, "test unknown error: details", erk.Message)
}

func TestErrorUserNotFound(t *testing.T) {
	erk := example1.ErrorUserNotFound("user %d not found", 123)
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(404), erk.Code)
	require.Equal(t, example1.ErrorReason_USER_NOT_FOUND.String(), erk.Reason)
	require.Equal(t, "user 123 not found", erk.Message)
}

func TestErrorInvalidParameter(t *testing.T) {
	erk := example1.ErrorInvalidParameter("parameter %s is invalid", "email")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(400), erk.Code)
	require.Equal(t, example1.ErrorReason_INVALID_PARAMETER.String(), erk.Reason)
	require.Equal(t, "parameter email is invalid", erk.Message)
}

func TestIsUnknown(t *testing.T) {
	erk := example1.ErrorUnknown("test error")
	t.Log(erk)
	require.True(t, example1.IsUnknown(erk))

	userErk := example1.ErrorUserNotFound("user not found")
	t.Log(userErk)
	require.False(t, example1.IsUnknown(userErk))

	stdErk := errors.New(500, "UNKNOWN", "standard error")
	t.Log(stdErk)
	require.True(t, example1.IsUnknown(stdErk))

	require.False(t, example1.IsUnknown(nil))
}

func TestIsUserNotFound(t *testing.T) {
	erk := example1.ErrorUserNotFound("user not found")
	t.Log(erk)
	require.True(t, example1.IsUserNotFound(erk))

	unknownErk := example1.ErrorUnknown("unknown error")
	t.Log(unknownErk)
	require.False(t, example1.IsUserNotFound(unknownErk))

	stdErk := errors.New(404, "USER_NOT_FOUND", "user not found")
	t.Log(stdErk)
	require.True(t, example1.IsUserNotFound(stdErk))
}

func TestIsInvalidParameter(t *testing.T) {
	erk := example1.ErrorInvalidParameter("invalid param")
	t.Log(erk)
	require.True(t, example1.IsInvalidParameter(erk))

	unknownErk := example1.ErrorUnknown("unknown error")
	t.Log(unknownErk)
	require.False(t, example1.IsInvalidParameter(unknownErk))

	stdErk := errors.New(400, "INVALID_PARAMETER", "invalid parameter")
	t.Log(stdErk)
	require.True(t, example1.IsInvalidParameter(stdErk))
}

func TestMetadataFieldName(t *testing.T) {
	erk := example1.ErrorUnknown("test error with metadata")
	t.Log(erk)

	require.NotNil(t, erk.Metadata)

	reason, exists := erk.Metadata["reason_numeric"]
	require.True(t, exists)

	expectedReason := fmt.Sprintf("%d", example1.ErrorReason_UNKNOWN.Number())
	require.Equal(t, expectedReason, reason)
}
