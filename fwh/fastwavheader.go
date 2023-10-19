package fwh

import "unsafe"

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
	RiffFileDescriptionHeader := file[0:4]
	WawDescriptionHeader := file[8:12]
	FmtDescriptionHeader := file[12:16]
	return WavInfo{
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

func (fwh *FastWavHeader) GetOnlyData(file []byte) []byte {
	return file[44:]
}
