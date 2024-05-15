package main

import (
	"encoding/base64"
	"flag"
	"io"
	"os"
)

func main() {
	decodeMode := flag.Bool("d", false, "decode input")
	_ = decodeMode
	help := flag.Bool("h", false, "display this message")
	inputFileName := flag.String("i", "-", `input file "-" for stdin`)
	_ = inputFileName
	outputFileName := flag.String("o", "-", `output file "-" for stdout`)
	_ = outputFileName

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	input := os.Stdin
	if *inputFileName != "-" {
		var err error
		input, err = os.Open(*inputFileName)
		if err != nil {
			panic(err)
		}
	}
	defer input.Close()

	output := os.Stdout
	if *outputFileName != "-" {
		var err error
		output, err = os.Create(*outputFileName)
		if err != nil {
			panic(err)
		}
	}
	defer output.Close()

	if *decodeMode {
		decode(input, output)
	} else {
		encode(input, output)
	}

	os.Exit(0)
}

func decode(input, output *os.File) {
	inputData, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}
	outputData, err := base64.RawURLEncoding.DecodeString(string(inputData))
	if err != nil {
		panic(err)
	}
	_, err = output.Write(outputData)
	if err != nil {
		panic(err)
	}
}
func encode(input, output *os.File) {
	inputData, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}
	outputData := base64.RawURLEncoding.EncodeToString(inputData)
	_, err = output.WriteString(outputData)
	if err != nil {
		panic(err)
	}
}
