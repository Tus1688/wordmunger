package main

import (
	"flag"
	"fmt"

	"github.com/Tus1688/wordmunger/pkg/wordmunger"
)

func main() {
	flag.Usage = func() {
		println("Usage: -f <input file> -o <output file> -c <concurrent>")
	}
	inputFile := flag.String("f", "", "input file")
	outputFile := flag.String("o", "", "output file")
	concurrent := flag.Int("c", 20, "concurrent")
	flag.Parse()
	if *inputFile == "" || *outputFile == "" {
		flag.Usage()
		return
	}
	arg := wordmunger.WordMunger{
		InputFile:  *inputFile,
		OutputFile: *outputFile,
		Concurrent: *concurrent,
	}
	arg.ReadFile()
	res := arg.Munging()
	fmt.Println(res)
	fmt.Println(len(res))
	err := arg.WriteFile(res)
	if err != nil {
		fmt.Println(err)
	}
}
