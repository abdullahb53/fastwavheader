package main

import (
	"log"
	"os"

	fwh "github.com/abdullahb5355/fastwavheader/fwh"
)

func main() {
	// Read a file.
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		log.Fatalf("File read err:%+x", err)
	}

	// Create instance.
	fwh := fwh.NewFastWavHeader()

	// Send data through FastWavHeader and get wav header info.
	myWavInfo := fwh.GetHeader(file)

	log.Println(myWavInfo)
}
