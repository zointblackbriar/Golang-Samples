package basic_http

import (
	"testing"
)

func MinimumComparison(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func TestMinimumComparison(t *testing.T) {
	ans := MinimumComparison(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2 ", ans)
	}
}
