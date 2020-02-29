package shouter

import (
	"bytes"
	"io"
	"testing"
)

func TestShouter(t *testing.T){
	var table = []struct {
		input string
		output string
	}{
		{input: "hello", output: "HELLO"},
		{input: "HellO", output: "HELLO"},
	}

	for _, row := range table {
		if v := Shout(row.input); v != row.output {
			t.Errorf("Expected %s but got %s\n", row.output, v)
		}  
	}
}

func TestShoutFile(t *testing.T){
	var table = []struct{
		input io.Reader
		output string
	}{
		{input: bytes.NewReader([]byte("Hello")), output: "HELLO"},
	}

	for _, row := range table {
		if v, _ := ShoutFile(row.input); v != row.output {
			t.Errorf("Should have returned %s rather than %s\n", row.output, v)
		}
	}

}