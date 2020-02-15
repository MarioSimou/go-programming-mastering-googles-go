package main

import (
	"fmt"

	p "./poetry"
)

func main() {
	var p = p.Poem{
		p.Stanza{
			"first line",
			"second line",
		},
		p.Stanza{
			"third line",
		},
	}
	var nv, nc = p.Statistics()
	fmt.Printf("nv: %d\tnc: %d\nnLines: %d\nnStanza: %d", nv, nc, p.NLines(), p.NStanza())
}
