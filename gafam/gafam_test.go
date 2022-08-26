package gafam

import "testing"

func TestLoadGafamData(t *testing.T) {
	var result = loadGafamData()
	if len(result) != 5 {
		t.Errorf("loadGafamData() = Got %d expected 5", len(result))
	}
}
