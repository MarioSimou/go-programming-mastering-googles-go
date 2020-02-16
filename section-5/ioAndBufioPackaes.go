package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type Line string
type Stanza []Line
type Poem []Stanza

func (p Poem) String() string {
	var lns []string
	for _, s := range p {
		for _, l := range s {
			lns = append(lns, string(l))
		}
	}
	return strings.Join(lns, "\n")
}

func main() {
	var name = "poem.txt"
	f, e := os.Open(name)
	if e != nil {
		log.Fatalln(e)
	}
	defer f.Close()

	var p Poem
	var s Stanza
	var scan = bufio.NewScanner(f)
	for scan.Scan() {
		var l = Line(scan.Text())
		if v, _ := utf8.DecodeRuneInString(string(l)); v == 65533 {
			p = append(p, s)
			s = Stanza{}
			continue
		}
		s = append(s, l)

	}
	p = append(p, s)

	if e := scan.Err(); e != nil {
		log.Fatalln(e)
	}

	fmt.Printf("%s", p)
}
