package main

import (
	"flag"
	"fmt"
	"io"
	"net.physalis/mofu/lib"
	"os"
	"strings"
)

func main() {
	var writers = []mofu.Writer{
		mofu.JavaProperties{},
		mofu.IOSStrings{},
		mofu.AndroidStrings{},
	}

	var identifiers []string
	for _, w := range writers {
		for _, ids := range w.Identifiers() {
			identifiers = append(identifiers, ids)
		}
	}

	var inputFile = flag.String("i", "", "Input file")
	var outputFile = flag.String("o", "", "Output file")
	var outputFormat = flag.String("f", "", "(optional) Output format ["+strings.Join(identifiers, ", ")+"]")
	flag.Parse()

	if len(*inputFile) < 1 || len(*outputFile) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	var writer mofu.Writer
	if len(*outputFormat) > 0 {
		for _, w := range writers {
			for _, i := range w.Identifiers() {
				if i == *outputFormat {
					writer = w
				}
			}
		}
	}

	if writer == nil {
		for _, w := range writers {
			if w.AcceptFile(*outputFile) {
				writer = w
			}
		}
	}

	if writer == nil {
		fmt.Printf("Error: no suitable output format found.\n")
		os.Exit(4)
	}

	in, err := os.Open(*inputFile)
	defer in.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	out, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer out.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(3)
	}

	reader := new(mofu.Yaml)

	config, err := reader.Read(io.Reader(in))
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(2)
	}

	writer.Write(config, io.Writer(out))
}
