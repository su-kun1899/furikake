package furikake_test

import (
	"github.com/su-kun1899/furikake"
	"testing"
)

func TestToCsv(t *testing.T) {
	jsonVal := `
[
	{"name":"Messi", "position":"Forward"},
	{"name":"Coutinho", "position":"Midfielder"},
	{"name":"Pique", "position":"Defender"}
]
`

	header := []string{"name", "position"}
	actual, err := furikake.ToCsv(header, jsonVal)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	expected := `"name","position"
"Messi","Forward"
"Coutinho","Midfielder"
"Pique","Defender"`

	if actual != expected {
		t.Errorf("ToCsv returns %v, expected %v", actual, expected)
	}
}
