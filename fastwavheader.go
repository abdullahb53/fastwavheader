package main

type WavInfo struct {
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

type FastWavHeader struct{}

func NewFastWavHeader() *FastWavHeader {
	return &FastWavHeader{}
}

func (fwh *FastWavHeader) GetHeader(file []byte) WavInfo {
	var head [44]byte
	for i := 0; i < 44; i++ {
		head[i] = file[i]
	}
	res := WavInfo{
		RiffFileDescriptionHeader: "",
		SizeOfFile:                uint32(head[4]) | uint32(head[5])<<8 | uint32(head[6])<<16 | uint32(head[7])<<24,
		WavDescriptonHeader:       "",
		FmtDescriptionHeader:      "",
		SizeOfWavSectionChunk:     uint32(head[16]) | uint32(head[17])<<8 | uint32(head[18])<<16 | uint32(head[19])<<24,
		TypeFormat:                uint16(head[20]) | uint16(head[21])<<8,
		MonoStereoFlag:            uint16(head[22]) | uint16(head[23])<<8,
		SampleFrequency:           uint32(head[24]) | uint32(head[25])<<8 | uint32(head[26])<<16 | uint32(head[27])<<24,
		AudioDataRateBytesSec:     uint32(head[28]) | uint32(head[29])<<8 | uint32(head[30])<<16 | uint32(head[31])<<24,
		BlockAlignment:            uint16(head[32]) | uint16(head[33])<<8,
		BitsPerSample:             uint16(head[34]) | uint16(head[35])<<8,
		DataDescriptionHeader:     uint32(head[36]) | uint32(head[37])<<8 | uint32(head[38])<<16 | uint32(head[39])<<24,
		SizeOfDataChunk:           uint32(head[40]) | uint32(head[41])<<8 | uint32(head[42])<<16 | uint32(head[43])<<24,
	}

	slice := head[0:16]
	slice[4] = 0
	slice[5] = 0
	slice[6] = 0
	slice[7] = 0

	str := string(slice)
	res.RiffFileDescriptionHeader = str[0:4]
	res.WavDescriptonHeader = str[8:12]
	res.FmtDescriptionHeader = str[12:16]
	return res
}

func (fwh *FastWavHeader) GetOnlyData(file []byte) []byte {
	return file[44:]
}
