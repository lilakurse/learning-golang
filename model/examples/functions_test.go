package examples

import (
	"testing"
)

func TestMax(t *testing.T) {
	var max int
	max = Max(48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17)
	if max != 97 {
		t.Error("Expected 97,got", max)
	}

}
