package poetry

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem(s ...Stanza) *Poem {
	return &Poem{}
}

func (p Poem) NStanza() int {
	return len(p)
}
func (p Poem) NLines() int {
	var counter = 0
	for _, s := range p {
		counter += len(s)
	}
	return counter
}

func (p Poem) Statistics() (int, int, int) {
	nVowels := 0
	nConsonants := 0
	nPuncs := 0
	for _, s := range p {
		for _, l := range s {
			for _, c := range l {
				switch c {
				case 'a', 'e', 'i', 'o', 'u':
					nVowels += 1
				case ',', ' ', ';', '!':
					nPuncs += 1
				default:
					nConsonants += 1
				}
			}
		}
	}

	return nVowels, nConsonants, nPuncs
}
