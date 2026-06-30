package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemMenuTreeListReturnsEmptySliceForEmptyNodes(t *testing.T) {
	service := NewSystemMenu()

	routes := service.treeList(nil)

	assert.NotNil(t, routes)
	assert.Empty(t, routes)
}
