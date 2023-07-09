package main

import (
	"fmt"
	"io"
	"log"
)

type Bytereader struct {
	content string
	pos     int
}

func (m *Bytereader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.content) {
		n := copy(p, m.content[m.pos:m.pos+1])
		m.pos++
		return n, nil
	}

	return 0, io.EOF
}

func MyReader() {

	myByteReader := Bytereader{
		content: "My First Sentence",
	}

	body, err := io.ReadAll(&myByteReader)

	if err != nil {
		log.Fatalf("Failed to parse the response %v", err)
	}

	fmt.Printf("My Slow reader output: %s \n", body)

}
