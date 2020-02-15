package poetry

import (
	"testing"
)

var p = Poem{
	Stanza{
		"first line",
		"second line",
	},
	Stanza{
		"third line",
	},
}

func TestNStanza(t *testing.T) {
	var ep Poem
	if v := ep.NStanza(); v != 0 {
		t.Fatalf("Should have returned %d stanzas rather than %d\n", 0, v)
	}
	if v := p.NStanza(); v != 2 {
		t.Fatalf("Should have returned %d stanzas rather than %d\n", 2, v)
	}
}

func TestNLines(t *testing.T) {
	var ep Poem
	if v := ep.NLines(); v != 0 {
		t.Fatalf("Should have returned %d lines rather than %d\n", 0, v)
	}
	if v := p.NLines(); v != 3 {
		t.Fatalf("Should have returned %d lines rather than %d\n", 3, v)
	}
}

func TestStats(t *testing.T) {
	var ep Poem
	if v, c, p := ep.Statistics(); v > 0 || c > 0 || p > 0 {
		t.Fatalf("Should have returned %d vowels,%d consonents,%d puncs rather than %d, %d, %d respectively\n", 0, 0, 0, v, c, p)
	}

	ep = Poem{Stanza{"hello"}}
	if v, c, p := ep.Statistics(); v != 2 || c != 3 || p != 0 {
		t.Fatalf("Should have returned %d vowels,%d consonents,%d puncs rather than %d, %d, %d respectively\n", 2, 3, 0, v, c, p)
	}
	ep = Poem{Stanza{"hello, world!"}}
	if v, c, p := ep.Statistics(); v != 3 || c != 7 || p != 3 {
		t.Fatalf("Should have returned %d vowels,%d consonents,%d puncs rather than %d, %d, %d respectively\n", 3, 7, 3, v, c, p)
	}
}
