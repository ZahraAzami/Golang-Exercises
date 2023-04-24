package main

import (
	"fmt"
	"io"
	"strings"
)

type CustomWriter struct {
}

func (cw *CustomWriter) Write(p []byte) (int, error) {
	return fmt.Printf("<< %s", string(p))
}

type WordCounter int

func (W *WordCounter) Write(p []byte) (n int, err error) {
	Word := strings.Fields(string(p))
	*W += WordCounter(len(Word))
	return len(Word), nil
}

type CounterWriter struct {
	W     io.Writer
	Count int64
}

func NewCounterWrit(new io.Writer) CounterWriter {
	return CounterWriter{W: new, Count: 0}
}
func (Cw *CounterWriter) Write(p []byte) (num int, err error) {
	num, err = Cw.W.Write(p)
	if err != nil {
		return num, err
	}
	Cw.Count = Cw.Count + int64(num)
	return num, err
}
func main() {

	cw := CustomWriter{}
	CWriter := NewCounterWrit(&cw)
	fmt.Fprintf(&CWriter, "Hear is : %s\n", "zaniar")
	fmt.Fprintf(&CWriter, "Hear is : %s\n", "mohamad")
	fmt.Fprintf(&CWriter, "Hear is : %d\n", CWriter.Count)
	var WC WordCounter
	fmt.Fprintf(&WC, "This is class room in class for Go learning")
}
