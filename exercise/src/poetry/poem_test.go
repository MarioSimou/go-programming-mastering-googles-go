package poetry

import (
	"log"
	"testing"
)

func TestNLines(t *testing.T) {
	var table = []struct {
		input  Poem
		output int
	}{
		{input: Poem{}, output: 0},
		{input: Poem{{"hello"}}, output: 1},
		{input: Poem{{"hello world"}, {"i'm marios"}}, output: 2},
	}

	for _, row := range table {
		if nl := row.input.NLines(); nl != row.output {
			t.Errorf("Should have returned %d rather than %d\n", row.output, nl)
		}
	}
}

func TestNStanza(t *testing.T) {
	var table = []struct {
		input  Poem
		output int
	}{
		{input: Poem{}, output: 0},
		{input: Poem{{"hello"}}, output: 1},
		{input: Poem{{"hello"}, {"world"}, {"tester"}}, output: 3},
	}

	for _, row := range table {
		if i, o := row.input.NStanza(), row.output; i != o {
			t.Errorf("Should have returned %d rather than %d\n", o, i)
		}
	}
}

func TestNWords(t *testing.T) {
	var table = []struct {
		input  Poem
		output int
	}{
		{input: Poem{}, output: 0},
		{input: Poem{{"hello"}}, output: 1},
		{input: Poem{{"hello world"}, {"this is awesome"}}, output: 5},
	}

	for _, row := range table {
		if i, o := row.input.NWords(), row.output; i != o {
			t.Errorf("Should have returned %d rather than %d\n", o, i)
		}
	}
}

func TestString(t *testing.T) {
	var table = []struct {
		input  Poem
		output string
	}{
		{input: Poem{{"hello"}}, output: "hello"},
		{input: Poem{{"hello"}, {"world"}, {"!"}}, output: "hello\nworld\n!"},
	}
	for _, row := range table {
		if i, o := row.input.String(), row.output; i != o {
			t.Errorf("Should have returned %q rather than %q\n", o, i)
		}
	}
}

func TestLoadPoem(t *testing.T) {
	var table = []struct {
		filename string
		output   string
	}{
		{filename: "sample.txt", output: "Two sets\nof family stories,\none long and detailed,\nabout many centuries\nof island ancestors, all living\non the same tropical farm...\nThe other side of the family tells stories\nthat are brief and vague, about violence\nin the Ukraine, which Dad's parents\nhad to flee forever, leaving all their\nloved ones\nbehind."},
	}

	for _, row := range table {
		p, e := LoadPoem(row.filename)
		if e != nil {
			log.Fatalln(e)
		}
		if i, o := p.String(), row.output; i != o {
			t.Errorf("Should have returned %q rather than %q\n", o, i)
		}
	}
}
