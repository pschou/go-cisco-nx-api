package client

import (
	"testing"
)

func TestDuration(t *testing.T) {
	d := Duration(20)
	err := d.UnmarshalText([]byte("P7DT12H2M5S"))
	if err != nil {
		t.Fatalf("Failed parse Duration %v", err)
	}
	if uint64(d) != 648125e9 {
		t.Fatalf("Failed Duration parse test %d != 648125e9", d)
	}

	err = d.UnmarshalText([]byte("P1MT4M31S"))
	if err != nil {
		t.Fatalf("Failed parse Duration %v", err)
	}
	if uint64(d) != 2592271e9 {
		t.Fatalf("Failed Duration parse test %d != 2592271e9", d)
	}

	err = d.UnmarshalText([]byte("1w2d"))
	if err != nil {
		t.Fatalf("Failed parse Duration %v", err)
	}
	if uint64(d) != 777600e9 {
		t.Fatalf("Failed Duration parse test %d != 777600e9", d)
	}

	err = d.UnmarshalText([]byte("PT4M31S"))
	if err != nil {
		t.Fatalf("Failed parse Duration %v", err)
	}
	if uint64(d) != 271e9 {
		t.Fatalf("Failed Duration parse test %d != 271e9", d)
	}

	err = d.UnmarshalText([]byte("00:04:30"))
	if err != nil {
		t.Fatalf("Failed parse Duration %v", err)
	}
	if uint64(d) != 270e9 {
		t.Fatalf("Failed Duration parse test %d != 270e9", d)
	}

	if d.String() != "PT4M30S" {
		t.Fatalf("Failed Duration String() test %s != \"PT4M30S\"", d.String())
	}
}
