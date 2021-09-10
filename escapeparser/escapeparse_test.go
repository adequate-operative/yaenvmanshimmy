package escapeparser

import (
	"testing"
)

func TestTokenise(t *testing.T) {
	tdata := []string{
		`Line 1`,
		`Line2 $Env$`,
		`Line 3`,
		`$$ Hello`,
		`$$ Hello $$`,
		`Hello $$`,
		`$$ Hello $$ There`,
	}
	for _, s := range tdata {
		stuff, err := Tokenise(s)
		if err != nil {
			t.Log("Failure!", err)
		}
		t.Log(stuff)
	}
}
