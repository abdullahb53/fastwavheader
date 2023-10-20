package fwh

import (
	"log"
	"os"
	"unsafe"
)

type WavInfo struct {
	FilePath string
	Header   Header
	err      string
}

type Header struct {
	RiffFileDescriptionHeader string
	SizeOfFile                uint32
	WavDescriptonHeader       string
	FmtDescriptionHeader      string
	SizeOfWavSectionChunk     uint32
	TypeFormat                uint16
	MonoStereoFlag            uint16
	SampleFrequency           uint32
	AudioDataRateBytesSec     uint32
	BlockAlignment            uint16
	BitsPerSample             uint16
	DataDescriptionHeader     uint32
	SizeOfDataChunk           uint32
}

type FastWavHeader struct {
	FilePathCh      chan string
	HeaderCh        chan WavInfo
	pauseSender     chan struct{}
	startSender     chan struct{}
	isChannelActive bool
}

func NewFastWavInstance() *FastWavHeader {
	return &FastWavHeader{
		pauseSender:     make(chan struct{}),
		startSender:     make(chan struct{}),
		isChannelActive: false,
	}
}

func (fwh *FastWavHeader) GetHeader(file []byte) Header {
	RiffFileDescriptionHeader := file[0:4]
	WawDescriptionHeader := file[8:12]
	FmtDescriptionHeader := file[12:16]
	return Header{
		RiffFileDescriptionHeader: unsafe.String(unsafe.SliceData(RiffFileDescriptionHeader), len(RiffFileDescriptionHeader)),
		SizeOfFile:                uint32(file[4]) | uint32(file[5])<<8 | uint32(file[6])<<16 | uint32(file[7])<<24,
		WavDescriptonHeader:       unsafe.String(unsafe.SliceData(WawDescriptionHeader), len(WawDescriptionHeader)),
		FmtDescriptionHeader:      unsafe.String(unsafe.SliceData(FmtDescriptionHeader), len(FmtDescriptionHeader)),
		SizeOfWavSectionChunk:     uint32(file[16]) | uint32(file[17])<<8 | uint32(file[18])<<16 | uint32(file[19])<<24,
		TypeFormat:                uint16(file[20]) | uint16(file[21])<<8,
		MonoStereoFlag:            uint16(file[22]) | uint16(file[23])<<8,
		SampleFrequency:           uint32(file[24]) | uint32(file[25])<<8 | uint32(file[26])<<16 | uint32(file[27])<<24,
		AudioDataRateBytesSec:     uint32(file[28]) | uint32(file[29])<<8 | uint32(file[30])<<16 | uint32(file[31])<<24,
		BlockAlignment:            uint16(file[32]) | uint16(file[33])<<8,
		BitsPerSample:             uint16(file[34]) | uint16(file[35])<<8,
		DataDescriptionHeader:     uint32(file[36]) | uint32(file[37])<<8 | uint32(file[38])<<16 | uint32(file[39])<<24,
		SizeOfDataChunk:           uint32(file[40]) | uint32(file[41])<<8 | uint32(file[42])<<16 | uint32(file[43])<<24,
	}
}

// StartStreamEvent has two channels.
// 'chan string' is for sendnig file-paths to channel. 'chan WavInfo' is for getting calculated WavHeaders from channel.
func (fwh *FastWavHeader) StartStreamEvent() {

	fwh.FilePathCh = make(chan string, 100)
	fwh.HeaderCh = make(chan WavInfo, 100)

	go func() {
		var (
			filePath      string
			file          *os.File
			err           error
			bufferCounter int
		)
		WavInfo := WavInfo{}
		for {
			select {
			case filePath = <-fwh.FilePathCh:
				WavInfo.FilePath = filePath
				file, err = os.Open(filePath)
				if err != nil {
					// log.Printf("Open file error: %+v", err)
					WavInfo.err = err.Error()
					WavInfo.Header = Header{}
					fwh.HeaderCh <- WavInfo
					file.Close()
					continue
				}
				bufferCounter = 0
				data := make([]byte, 44)
				for {
					n, err := file.Read(data)
					if n == 0 || err != nil {
						break
					}
					bufferCounter += n
					if bufferCounter >= 44 {
						WavInfo.Header = fwh.GetHeader(data)
						WavInfo.err = ""
						fwh.HeaderCh <- WavInfo
						file.Close()
						break
					}
				}

			case <-fwh.pauseSender:
				<-fwh.startSender
			}
		}
	}()
}

// Change stream channels'(producer filepaths, consumer wavheaderinfo) queue size.
func (fwh *FastWavHeader) ChangeQueueSize(FilePathProducerSize int, WavInfoConsumerSize int) {
	oldProducer := fwh.FilePathCh
	oldProducerPause := fwh.pauseSender
	oldProducerStarter := fwh.startSender
	oldConsumer := fwh.HeaderCh

	CapOfProducer := cap(oldProducer)
	CapOfConsumer := cap(oldConsumer)

	oldProducerPause <- struct{}{}
	log.Println("[FaswWavHeader] Stream is gracefully shutdown.")

	fwh.HeaderCh = make(chan WavInfo, WavInfoConsumerSize)
	log.Printf("[FaswWavHeader] Consumer WavInfo queue size %v is changed to %v\n", CapOfConsumer, cap(fwh.HeaderCh))

	fwh.FilePathCh = make(chan string, FilePathProducerSize)
	log.Printf("[FaswWavHeader] Producer file path queue size %v is changed to %v\n", CapOfProducer, cap(fwh.FilePathCh))

	close(oldProducer)
	close(oldConsumer)
	oldProducerStarter <- struct{}{}
	log.Println("[FaswWavHeader] Stream is started..")

	temp := []string{}
	for val := range oldProducer {
		log.Println("[FastWavHeader] Secured >>", val)
		temp = append(temp, val)
	}

	// Send temp files to new channel.
	for _, val := range temp {
		fwh.FilePathCh <- val
	}
	log.Println("[FastWavHeader] ### Retrieving is completed ###")
}
