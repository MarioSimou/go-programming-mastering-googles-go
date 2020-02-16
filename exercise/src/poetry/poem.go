package poetry

import (
	"bufio"
	"os"
	"strings"
)

type Line string
type Stanza []Line
type Poem []Stanza

type Poetry interface {
	NLines() int
	NStanza() int
	NWords() int
	String() string
}

func (p Poem) NLines() int {
	var c = 0
	for _, s := range p {
		c += len(s)
	}
	return c
}

func (p Poem) NStanza() int {
	return len(p)
}

func (p Poem) NWords() int {
	var c = 0
	for _, s := range p {
		for _, l := range s {
			c += len(strings.Split(string(l), " "))
		}
	}
	return c
}

func (p Poem) String() string {
	var text []string
	for _, s := range p {
		for _, l := range s {
			text = append(text, string(l))
		}
	}
	return strings.Join(text, "\n")
}

func LoadPoem(fileName string) (Poetry, error) {
	f, e := os.Open(fileName)
	if e != nil {
		return nil, e
	}

	var s = bufio.NewScanner(f)
	var p Poem
	var stanza Stanza
	for s.Scan() {
		var ts = s.Text()
		if ts == "" {
			p = append(p, stanza)
			stanza = Stanza{}
			continue
		}
		stanza = append(stanza, Line(ts))
	}
	p = append(p, stanza)

	return p, nil
}
