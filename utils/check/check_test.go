package check

import (
	"github.com/stretchr/testify/assert"
	"richii_go/models"
	"testing"
)

var testCase = []models.HaiArr{
	{
		{2, 2, 0, 2, 0, 0, 2, 2, 2},
		{0, 0, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 2, 1, 1, 1, 1, 1, 0, 0},
	},
	{
		{0, 0, 0, 0, 0, 2, 2, 2, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 3, 0, 0, 0},
	},
	{
		{2, 2, 2, 2, 0, 0, 2, 2, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
}

func TestIf7D(t *testing.T) {
	assert.True(t, Is7D(testCase[0]))
	assert.False(t, Is7D(testCase[1]))
	assert.False(t, Is7D(testCase[2]))
	assert.True(t, Is7D(testCase[3]))
}

func TestIf13G(t *testing.T) {
	assert.False(t, Is13G(testCase[0]))
	assert.True(t, Is13G(testCase[1]))
	assert.False(t, Is13G(testCase[0]))
	assert.True(t, Is13G(testCase[1]))
}

func TestIsNormal(t *testing.T) {
	assert.False(t, IsNormal(testCase[0]))
	assert.False(t, IsNormal(testCase[1]))
	assert.True(t, IsNormal(testCase[2]))
	assert.True(t, IsNormal(testCase[3]))
}

func TestIsValid(t *testing.T) {
	assert.True(t, IsValid(testCase[0]))
	assert.True(t, IsValid(testCase[1]))
	assert.True(t, IsValid(testCase[2]))
	assert.True(t, IsValid(testCase[3]))
}
