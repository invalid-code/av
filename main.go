package main

import (
	"bufio"
	"os"
)

func main() {
//	signature-based detection
	signaturesFile, err := os.Open("signatures.csv")
	if err != nil {
		panic(err)
	}
	signatureFileReader := bufio.NewReader(signaturesFile)
	signatureFileReader.ReadLine()
//	heuristic-based detection
}