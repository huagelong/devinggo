package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapterGetTable(t *testing.T) {
	adapter := Adapter{}
	ctx := context.Background()

	testCases := []struct {
		name     string
		key      interface{}
		expected string
	}{
		{
			name:     "valid select cache key",
			key:      "SelectCache:system_user@abc123",
			expected: "system_user",
		},
		{
			name:     "valid select cache key with complex table name",
			key:      "SelectCache:system_user_role@def456",
			expected: "system_user_role",
		},
		{
			name:     "non select cache key",
			key:      "user:123",
			expected: "",
		},
		{
			name:     "empty key",
			key:      "",
			expected: "",
		},
		{
			name:     "select cache prefix only",
			key:      "SelectCache:",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := adapter.getTable(ctx, tc.key)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetTagCacheInstance(t *testing.T) {
	oldTagCache := _tagCache
	_tagCache = nil
	defer func() { _tagCache = oldTagCache }()

	_, err := getTagCacheInstance()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tag cache uninitialized")
}

func TestGetCacheUninitialized(t *testing.T) {
	oldCache := _cache
	_cache = nil
	defer func() { _cache = oldCache }()

	_, err := GetCache()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cache uninitialized")
}
