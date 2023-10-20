package main

import (
	"log"
	"os"

	fwh "github.com/abdullahb53/fastwavheader/fwh"
)

func main() {
	// // Read a file.
	file, err := os.ReadFile("./test/getheader/sounds/ImperialMarch60.wav")
	if err != nil {
		log.Fatalf("File read err:%+x", err)
	}
	// Create instance.
	fwh := fwh.NewFastWavInstance()
	// // Send data through FastWavHeader and get wav header info.
	myWavInfo := fwh.GetHeader(file)
	log.Println(myWavInfo)

	// If u want to stream data through channels.
	fwh.StartStreamEvent()
	filePaths := []string{
		"./test/getheader/sounds/ImperialMarch60.wav",
		"./test/getheader/sounds/ImperialMarch61.wav",
		"./test/getheader/sounds/ImperialMarch62.wav",
		"./test/getheader/sounds/ImperialMarch63.wav",
		"./test/getheader/sounds/ImperialMarch64.wav",
		"./test/getheader/sounds/ImperialMarch65.wav",
		"./test/getheader/sounds/ImperialMarch66.wav",
		"./test/getheader/sounds/ImperialMarch67.wav",
		"./test/getheader/sounds/ImperialMarch68.wav",
		"./test/getheader/sounds/ImperialMarch69.wav",
		"./test/getheader/sounds/ImperialMarch70.wav",
		"./test/getheader/sounds/ImperialMarch71.wav",
	}

	filePaths2 := []string{
		"./test/getheader/sounds/ImperialMarch72.wav",
		"./test/getheader/sounds/ImperialMarch73.wav",
		"./test/getheader/sounds/ImperialMarch74.wav",
		"./test/getheader/sounds/ImperialMarch75.wav",
		"./test/getheader/sounds/ImperialMarch76.wav",
		"./test/getheader/sounds/ImperialMarch77.wav",
	}

	// Send your filePaths to channel.
	for _, val := range filePaths {
		fwh.FilePathCh <- val
	}

	// Consume WavHeaderInfos from channel.
	go func() {
		consume := fwh.HeaderCh
		for {
			select {
			case wavinfo, ok := <-consume:
				if !ok {
					consume = fwh.HeaderCh
				} else {
					log.Println("@@ WavInfo:", wavinfo)
				}
			default:
			}
		}
	}()

	// Change channel capacity. Async or Sync.
	go fwh.ChangeQueueSize(30, 40)

	// Send your filePaths to adjusted channel.
	for _, val := range filePaths2 {
		fwh.FilePathCh <- val
	}

	<-make(chan struct{})

}
