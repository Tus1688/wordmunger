package main

import "flag"

func main() {
	flag.Usage = func() {
		println("Usage: -f <input file> -o <output file>")
	}
	flag.Parse()
	inputFile := flag.String("f", "", "input file")
	outputFile := flag.String("o", "", "output file")
	if *inputFile == "" || *outputFile == "" {
		flag.Usage()
		return
	}
}
