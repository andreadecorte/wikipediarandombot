package apis

import (
	"testing"
)

// TestReadingTime checks the computation of reading time
func TestReadingTime(t *testing.T) {
	minutes := timeToRead(100)
	if minutes != 1 {
		t.Fatalf("We should return 1, not %d minutes", minutes)
	}
}

// TestReadingTime2 checks the computation of reading time
func TestReadingTime2(t *testing.T) {
	minutes := timeToRead(900)
	if minutes != 3 {
		t.Fatalf("We should return 3, not %d minutes", minutes)
	}
}

