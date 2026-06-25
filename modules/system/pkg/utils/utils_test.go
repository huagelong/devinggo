// Package utils
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil error", nil, false},
		{"sql.ErrNoRows", sql.ErrNoRows, false},
		{"normal error", fmt.Errorf("some error"), true},
		{"wrapped sql.ErrNoRows", fmt.Errorf("wrapped: %w", sql.ErrNoRows), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsError(tc.err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetQueryMap(t *testing.T) {
	testCases := []struct {
		name        string
		query       string
		expectError bool
		expected    map[string]interface{}
	}{
		{
			name:        "valid query string",
			query:       "name=john&age=25&active=true",
			expectError: false,
			expected: map[string]interface{}{
				"name":   "john",
				"age":    "25",
				"active": "true",
			},
		},
		{
			name:        "empty query string",
			query:       "",
			expectError: false,
			expected:    map[string]interface{}{},
		},
		{
			name:        "query with special characters",
			query:       "key=value%20with%20spaces",
			expectError: false,
			expected: map[string]interface{}{
				"key": "value with spaces",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetQueryMap(tc.query)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				for k, v := range tc.expected {
					assert.Equal(t, v, result[k])
				}
			}
		})
	}
}

func TestGetConnectDbName(t *testing.T) {
	testCases := []struct {
		name        string
		dsn         string
		expectError bool
		expected    string
	}{
		{
			name:        "valid postgres dsn",
			dsn:         "user:password@tcp(localhost:5432)/mydb",
			expectError: false,
			expected:    "mydb",
		},
		{
			name:        "valid mysql dsn",
			dsn:         "root:pass@tcp(127.0.0.1:3306)/testdb",
			expectError: false,
			expected:    "testdb",
		},
		{
			name:        "invalid dsn format",
			dsn:         "invalid-dsn",
			expectError: true,
			expected:    "",
		},
		{
			name:        "empty dsn",
			dsn:         "",
			expectError: true,
			expected:    "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetConnectDbName(tc.dsn)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestGetModule(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected string
	}{
		{"normal path", "/api/user/list", "api"},
		{"root path", "/", "system"},
		{"empty path", "", "system"},
		{"single slash", "/", "system"},
		{"path with multiple segments", "/system/user/profile", "system"},
		{"path with empty second segment", "//user", "system"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetModule(tc.path)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFileMd5(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	testContent := "Hello, World!"

	err := os.WriteFile(testFile, []byte(testContent), 0644)
	assert.NoError(t, err)

	hash, err := FileMd5(testFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.Equal(t, "65a8e27d8879283831b664bd8b7f0ad4", hash)

	_, err = FileMd5(filepath.Join(tempDir, "nonexistent.txt"))
	assert.Error(t, err)
}

func TestMergeAndDeduplicateWithFunc(t *testing.T) {
	type Item struct {
		ID   int
		Name string
	}

	testCases := []struct {
		name        string
		compareFunc func(Item) string
		lists       [][]Item
		expected    []Item
	}{
		{
			name: "merge with duplicates",
			compareFunc: func(i Item) string {
				return fmt.Sprintf("%d", i.ID)
			},
			lists: [][]Item{
				{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}},
				{{ID: 2, Name: "B"}, {ID: 3, Name: "C"}},
			},
			expected: []Item{
				{ID: 1, Name: "A"},
				{ID: 2, Name: "B"},
				{ID: 3, Name: "C"},
			},
		},
		{
			name: "empty lists",
			compareFunc: func(i Item) string {
				return fmt.Sprintf("%d", i.ID)
			},
			lists:    [][]Item{},
			expected: nil,
		},
		{
			name: "single list",
			compareFunc: func(i Item) string {
				return fmt.Sprintf("%d", i.ID)
			},
			lists: [][]Item{
				{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}},
			},
			expected: []Item{
				{ID: 1, Name: "A"},
				{ID: 2, Name: "B"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MergeAndDeduplicateWithFunc(tc.compareFunc, tc.lists...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestReplaceSubstr(t *testing.T) {
	testCases := []struct {
		name      string
		s         string
		oldSubstr string
		newSubstr string
		expected  string
	}{
		{"replace single occurrence", "hello world", "world", "Go", "hello Go"},
		{"replace multiple occurrences", "aaa bbb aaa", "aaa", "ccc", "ccc bbb ccc"},
		{"no match", "hello world", "xyz", "abc", "hello world"},
		{"empty string", "", "old", "new", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReplaceSubstr(tc.s, tc.oldSubstr, tc.newSubstr)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestHasField(t *testing.T) {
	type TestStruct struct {
		Name  string
		Age   int
		Email string
	}

	testCases := []struct {
		name      string
		obj       interface{}
		fieldName string
		expected  bool
	}{
		{"existing field", TestStruct{Name: "John"}, "Name", true},
		{"non-existing field", TestStruct{Name: "John"}, "Address", false},
		{"pointer to struct", &TestStruct{Name: "John"}, "Name", true},
		{"non-struct type", "string", "Length", false},
		{"empty field name", TestStruct{Name: "John"}, "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := HasField(tc.obj, tc.fieldName)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetRootPath(t *testing.T) {
	path := GetRootPath()
	assert.NotEmpty(t, path)
}

func TestGetTmpDir(t *testing.T) {
	dir := GetTmpDir()
	assert.NotEmpty(t, dir)
}

func TestGetFieldQuote(t *testing.T) {
	quote := GetFieldQuote()
	assert.Equal(t, "\"", quote)
}

func TestQuoteField(t *testing.T) {
	testCases := []struct {
		name      string
		fieldName string
		expected  string
	}{
		{"simple field", "name", "\"name\""},
		{"field with spaces", "user name", "\"user name\""},
		{"empty field", "", "\"\""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := QuoteField(tc.fieldName)
			assert.Equal(t, tc.expected, result)
		})
	}
}
