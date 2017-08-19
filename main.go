package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/akr4/mofu/lib"
)

func main() {
	var writers = []mofu.Writer{
		mofu.JavaProperties{},
		mofu.IOSStrings{},
		mofu.AndroidStrings{},
		mofu.Json{},
	}

	var identifiers []string
	for _, w := range writers {
		for _, ids := range w.Identifiers() {
			identifiers = append(identifiers, ids)
		}
	}

	var inputFile = flag.String("i", "", "Input file")
	var outputFile = flag.String("o", "", "(optional) Output file")
	var outputFormat = flag.String("f", "", "(optional) Output format ["+strings.Join(identifiers, ", ")+"]")
	var includes = flag.String("includes", "", "(optional) Output strings only for items which have the prefix.")
	flag.Parse()

	if len(*inputFile) < 1 {
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

	if writer == nil && len(*outputFile) > 0 {
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

	reader := new(mofu.Yaml)

	config, err := reader.Read(io.Reader(in))
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(2)
	}

	if len(*includes) > 0 {
		config = filter(config, *includes)
	}

	out, err := os.Stdout, error(nil)
	if len(*outputFile) > 0 {
		out, err = os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	}
	defer out.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(3)
	}

	writer.Write(config, io.Writer(out))
}

func filter(config *mofu.Config, includes string) *mofu.Config {
	var configs []mofu.Config
	for _, include := range strings.Split(includes, ",") {
		var key mofu.Key
		for _, token := range strings.Split(include, ".") {
			key = append(key, token)
		}
		configs = append(configs, config.Filter(key))
	}
	var newConfig mofu.Config
	for _, c := range configs {
		newConfig = newConfig.Merge(c)
	}
	return &newConfig
}
