package ciscobcs

import (
	"os"
	"strings"
	"testing"
)

func TestBulk(t *testing.T) {
	file, err := os.ReadFile("../../demo_bcs_bulk.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("bulk scanner linecount", func(t *testing.T) {
		want := BulkResults{
			LineCount: 996,
		}
		got, err := scanBulk(strings.NewReader(string(file)))
		if err != nil {
			t.Errorf("didn't expect error reading scanning file")
		}
		if got.LineCount != want.LineCount {
			t.Errorf("got %v; want %v", got.LineCount, want.LineCount)
		}
	})
	t.Run("bulk scanner typecount", func(t *testing.T) {
		countOfTypesWanted := make(map[string]int)
		countOfTypesWanted["fn_bulletin"] = 103
		countOfTypesWanted["psirt_bulletin"] = 353
		countOfTypesWanted["device"] = 300
		countOfTypesWanted["track_summary"] = 20
		countOfTypesWanted["track_smupie_recommendation"] = 2
		countOfTypesWanted["sw_eox_bulletin"] = 34
		countOfTypesWanted["hw_eox_bulletin"] = 184

		got, err := scanBulk(strings.NewReader(string(file)))
		if err != nil {
			t.Errorf("didn't expect error reading scanning file")
		}
		for wantkey, wantval := range countOfTypesWanted {
			if gotval, ok := got.CountOfTypes[wantkey]; ok {
				if gotval != wantval {
					t.Errorf("%v: got %v; want %v", wantkey, gotval, wantval)
				}
			} else {
				t.Errorf("missing key %v in countOfTypes", wantkey)
			}
		}
	})
}

func BenchmarkBulk(b *testing.B) {
	file, err := os.ReadFile("../../demo_bcs_bulk.jsonl")
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		filereader := strings.NewReader(string(file))
		scanBulk(filereader)
	}
}
