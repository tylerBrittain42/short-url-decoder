package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tylerBrittain42/short-url-decoder/pkg/decoder"
)

func main() {
	urlPtr := flag.String("url", "", "the url to be decoded")
	outputPtr := flag.String("output", "final", "specify the output type(final, trace, csv)")

	flag.Parse()

	// should not function if url is empty OR both trace and csv are specified
	// OK if neither are, use default(final product)
	if *urlPtr == "" {
		fmt.Print("A url value must be provided\n")
		os.Exit(1)
	}
	switch {
	case *outputPtr == "trace":
		list, err := decoder.Trace(*urlPtr)
		if err != nil {
			fmt.Printf("Error: %w\n")
			os.Exit(1)
		}
		fmt.Println(strings.Join(list, " -> "))

	case *outputPtr == "csv":
		list, err := decoder.Trace(*urlPtr)
		if err != nil {
			fmt.Printf("Error: %w\n")
			os.Exit(1)
		}
		fmt.Println(strings.Join(list, ","))
	default:
		res, err := decoder.FinalDestination(*urlPtr)
		if err != nil {
			fmt.Printf("Error: %w\n")
			os.Exit(1)
		}
		fmt.Println(res)
	}

}
