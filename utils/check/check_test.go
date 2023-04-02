package check

import (
	"github.com/stretchr/testify/assert"
	"richii_go/models"
	"testing"
)

var testCase1 models.PaiArr
var testCase2 models.PaiArr
var testCase3 models.PaiArr
var testCase4 models.PaiArr

func TestCheck7D(t *testing.T) {
	testCase1 = models.PaiArr{
		M:    [9]int{2, 2, 0, 2, 0, 0, 2, 2, 2},
		P:    [9]int{0, 0, 2, 0, 0, 0, 0, 0, 0},
		S:    [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		Z:    [7]int{0, 0, 0, 0, 0, 0, 0},
		Dora: 0,
	}
	assert.True(t, Check7D(testCase1))
}
