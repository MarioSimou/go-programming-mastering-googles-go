package shouter

import (
	"bufio"
	"io"
	"strings"
)

func Shout(s string) string {
	return strings.ToUpper(s)
}

func ShoutLoud(s string) string {
	return strings.ToLower(s)
}

func ShoutFile(f io.Reader) (string, error) {
		var scanner = bufio.NewScanner(f)
		var text []string

		for scanner.Scan() {
			var t = scanner.Text()
			if e := scanner.Err(); e != nil {
				return "", e
			}
			text = append(text, Shout(t))
		}

		return strings.Join(text,"\n"), nil
}


// Read(p []byte) (n int, err error)