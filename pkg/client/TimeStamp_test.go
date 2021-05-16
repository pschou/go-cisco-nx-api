package client

import (
	"testing"
	//"time"
)

func TestTime(t *testing.T) {
	d := TimeStamp(0)
	err := d.UnmarshalText([]byte("11/04/2019 22:13:33"))
	if err != nil {
		t.Fatalf("Failed parse Time %v", err)
	}
	val := d.Time()
	if val.Unix() != 1572905613 {
		t.Fatalf("Failed TimeStamp parse test %v != 1572905613", val.Unix())
	}
	if d.String() != "11/04/2019 22:13:33" {
		t.Fatalf("Failed TimeStamp parse test %s != \"11/04/2019 22:13:33\"", d.String())
	}

	d, err = ParseTimeStamp("11/04/2019 22:13:32")
	if err != nil {
		t.Fatalf("Failed parse Time %v", err)
	}
	if d.String() != "11/04/2019 22:13:32" {
		t.Fatalf("Failed TimeStamp parse test %s != \"11/04/2019 22:13:32\"", d.String())
	}

	d, err = ParseTimeStamp("11/04/2019")
	if err != nil {
		t.Fatalf("Failed parse Time %v", err)
	}
	if d.String() != "11/04/2019 00:00:00" {
		t.Fatalf("Failed TimeStamp parse test %s != \"11/04/2019 00:00:00\"", d.String())
	}
}
