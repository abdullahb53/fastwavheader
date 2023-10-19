package main

import (
	"os"
	"testing"
	"unsafe"
)

func FastWavHeaderSlice(file []byte) WavInfo {
	return WavInfo{
		RiffFileDescriptionHeader: string(file[0:4]),
		SizeOfFile:                uint32(file[4]) | uint32(file[5])<<8 | uint32(file[6])<<16 | uint32(file[7])<<24,
		WavDescriptonHeader:       string(file[8:12]),
		FmtDescriptionHeader:      string(file[12:16]),
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

func FastWavHeaderSliceUnsafe(file []byte) WavInfo {
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

func FastWavHeaderSliceReflect(file []byte) WavInfo {
	RiffFileDescriptionHeader := file[0:4]
	WawDescriptionHeader := file[8:12]
	FmtDescriptionHeader := file[12:16]
	return WavInfo{
		RiffFileDescriptionHeader: *(*string)(unsafe.Pointer(&RiffFileDescriptionHeader)),
		SizeOfFile:                uint32(file[4]) | uint32(file[5])<<8 | uint32(file[6])<<16 | uint32(file[7])<<24,
		WavDescriptonHeader:       *(*string)(unsafe.Pointer(&WawDescriptionHeader)),
		FmtDescriptionHeader:      *(*string)(unsafe.Pointer(&FmtDescriptionHeader)),
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

func FastWavHeaderSliceFancy(file []byte) WavInfo {
	res := WavInfo{
		RiffFileDescriptionHeader: "",
		SizeOfFile:                uint32(file[4]) | uint32(file[5])<<8 | uint32(file[6])<<16 | uint32(file[7])<<24,
		WavDescriptonHeader:       "",
		FmtDescriptionHeader:      "",
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

	slice := file[0:16]
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

func FastWavHeaderArrayUnsafe(file []byte) WavInfo {
	var (
		head [44]byte
	)
	for i := 0; i < 44; i++ {
		head[i] = file[i]

	}
	RiffFileDescriptionHeader := head[0:4]
	WawDescriptionHeader := head[8:12]
	FmtDescriptionHeader := head[12:16]
	return WavInfo{
		RiffFileDescriptionHeader: unsafe.String(unsafe.SliceData(RiffFileDescriptionHeader), len(RiffFileDescriptionHeader)),
		SizeOfFile:                uint32(head[4]) | uint32(head[5])<<8 | uint32(head[6])<<16 | uint32(head[7])<<24,
		WavDescriptonHeader:       unsafe.String(unsafe.SliceData(WawDescriptionHeader), len(WawDescriptionHeader)),
		FmtDescriptionHeader:      unsafe.String(unsafe.SliceData(FmtDescriptionHeader), len(FmtDescriptionHeader)),
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
}

func FastWavHeaderArrayReflect(file []byte) WavInfo {
	var (
		head [44]byte
	)
	for i := 0; i < 44; i++ {
		head[i] = file[i]
	}
	RiffFileDescriptionHeader := head[0:4]
	WawDescriptionHeader := head[8:12]
	FmtDescriptionHeader := head[12:16]
	return WavInfo{
		RiffFileDescriptionHeader: *(*string)(unsafe.Pointer(&RiffFileDescriptionHeader)),
		SizeOfFile:                uint32(head[4]) | uint32(head[5])<<8 | uint32(head[6])<<16 | uint32(head[7])<<24,
		WavDescriptonHeader:       *(*string)(unsafe.Pointer(&WawDescriptionHeader)),
		FmtDescriptionHeader:      *(*string)(unsafe.Pointer(&FmtDescriptionHeader)),
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
}

func FastWavHeaderArrayFancy(file []byte) WavInfo {
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

func BenchmarkWithSlice(b *testing.B) {
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		b.Fatalf("[BENCH] read err: %+x", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastWavHeaderSlice(file)
	}
}

func BenchmarkWithSliceUnsafe(b *testing.B) {
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		b.Fatalf("[BENCH] read err: %+x", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastWavHeaderSliceUnsafe(file)
	}
}
func BenchmarkWithSliceReflect(b *testing.B) {
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		b.Fatalf("[BENCH] read err: %+x", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastWavHeaderSliceReflect(file)
	}
}

func BenchmarkWithSliceFancy(b *testing.B) {
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		b.Fatalf("[BENCH] read err: %+x", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastWavHeaderSliceFancy(file)
	}
}

func BenchmarkWithArrayUnsafe(b *testing.B) {
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		b.Fatalf("[BENCH] read err: %+x", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastWavHeaderArrayUnsafe(file)
	}
}

func BenchmarkWithArrayReflect(b *testing.B) {
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		b.Fatalf("[BENCH] read err: %+x", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastWavHeaderArrayReflect(file)
	}
}

func BenchmarkWithArrayFancy(b *testing.B) {
	file, err := os.ReadFile("./sounds/ImperialMarch60.wav")
	if err != nil {
		b.Fatalf("[BENCH] read err: %+x", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastWavHeaderArrayFancy(file)
	}
}
