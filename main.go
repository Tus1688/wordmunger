package main

import (
	"flag"
	"fmt"

	"github.com/Tus1688/wordmunger/pkg/wordmunger"
)

func main() {
	flag.Usage = func() {
		println("Usage: -f <input file> -o <output file>")
	}
	inputFile := flag.String("f", "", "input file")
	outputFile := flag.String("o", "", "output file")
	flag.Parse()
	if *inputFile == "" || *outputFile == "" {
		flag.Usage()
		return
	}
	arg := wordmunger.WordMunger{
		InputFile:  *inputFile,
		OutputFile: *outputFile,
	}
	arg.ReadFile()
	res := arg.Munging()
	// blue color
	fmt.Println("[i] generated:", len(res), "words")
	err := arg.WriteFile(res)
	if err != nil {
		fmt.Println("[-] error:", err)
	}
	fmt.Println("[*] done")
}
