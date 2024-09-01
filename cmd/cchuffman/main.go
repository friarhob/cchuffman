package main

import (
	"fmt"
	"os"

	"github.com/friarhob/cchuffman/internal/encoder"
	"github.com/friarhob/cchuffman/internal/exitcodes"
)

func printHelpMessage() {
	fmt.Fprintln(os.Stderr, "Usage: cchuffman input_filepath output_filepath")
}

func execute() {
	if len(os.Args) != 2 {
		printHelpMessage()
		os.Exit(int(exitcodes.UsageError))
	}

	if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		printHelpMessage()
		os.Exit(int(exitcodes.OK))
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s", os.Args[1])
		os.Exit(int(exitcodes.ErrorReadingFile))
	}

	defer file.Close()

	frequencySlice := encoder.FrequencySlice(file)

	for _, kvPair := range frequencySlice {
		fmt.Printf("%c: %d\n", kvPair[0].(rune), kvPair[1].(int))
	}
}

func main() {
	execute()
}


