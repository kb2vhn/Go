// io-example Black Hat Go book

package main

import(
	"fmt"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin
type FooReader struct{} // create the struct as empty

// Read reads data from stdin.
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ") // add "in > " with no extra formating
	return os.Stdin.Read(b)
}

// FooWriter defines an io.Writer to write to Stdout
type FooWriter struct{} // create the struct as empty

// Write writes data to Stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ") // same as reader just in reverse
	return os.Stdout.Write(b)
}

func main() {
	// Instantiate reader and writer.
	var (
		reader FooReader
		writer FooWriter
	)

	// Create buffer to hold input/output.
	input := make([]byte, 4096)
	// Use reader to read input.
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// User writer to write output.
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
