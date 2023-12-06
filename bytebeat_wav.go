// 06-12-2023 i4k
// This code creates an output.wav file. You can open it in any media player.

package main

import (
	"encoding/binary"
	"os"
)

func main() {
	const sampleRate = 8000
	const duration = 16 // in seconds

	file, err := os.Create("output.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// WAV header
	_ = binary.Write(file, binary.LittleEndian, []byte("RIFF"))
	_ = binary.Write(file, binary.LittleEndian, uint32(36+sampleRate*duration)) // File size - 36 bytes (header) + sound size
	_ = binary.Write(file, binary.LittleEndian, []byte("WAVE"))
	_ = binary.Write(file, binary.LittleEndian, []byte("fmt "))
	_ = binary.Write(file, binary.LittleEndian, uint32(16))
	_ = binary.Write(file, binary.LittleEndian, uint16(1))          // audioformat (1 for PCM)
	_ = binary.Write(file, binary.LittleEndian, uint16(1))          // Channels
	_ = binary.Write(file, binary.LittleEndian, uint32(sampleRate)) // Sampling frequency
	_ = binary.Write(file, binary.LittleEndian, uint32(sampleRate*1))
	_ = binary.Write(file, binary.LittleEndian, uint16(1))
	_ = binary.Write(file, binary.LittleEndian, uint16(8))
	_ = binary.Write(file, binary.LittleEndian, []byte("data"))
	_ = binary.Write(file, binary.LittleEndian, uint32(sampleRate*duration))

	// Генерация звука
	for t := 0; t < sampleRate*duration; t++ {
		val := uint8((t & int(t>>8)))
		_ = binary.Write(file, binary.LittleEndian, val)
	}

	println("Done! output.wav generated.")
}
