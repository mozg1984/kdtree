package kdtree

import (
	"testing"
)

func TestIndex_Within(t *testing.T) {
	points := []Point{{48.381124, 54.310801}}
	p := Point{48.396230, 54.312595}

	index := NewIndex()
	index.Build(points)

	found := index.Within(p, 1010)

	if len(found) == 0 {
		t.Error("A point within ~1 km was not found")
	}

	if len(found) == 1 && (found[0].Lon() != 48.381124 || found[0].Lat() != 54.310801) {
		t.Error("Incorrect point was found")
	}
}
