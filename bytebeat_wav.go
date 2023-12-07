// 07-12-2023 i4k
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
	binary.Write(file, binary.LittleEndian, []byte("RIFF"))
	binary.Write(file, binary.LittleEndian, uint32(36+sampleRate*duration)) // File size - 36 bytes (header) + sound size
	binary.Write(file, binary.LittleEndian, []byte("WAVE"))
	binary.Write(file, binary.LittleEndian, []byte("fmt "))
	binary.Write(file, binary.LittleEndian, uint32(16))
	binary.Write(file, binary.LittleEndian, uint16(1))          // audioformat (1 for PCM)
	binary.Write(file, binary.LittleEndian, uint16(1))          // Channels
	binary.Write(file, binary.LittleEndian, uint32(sampleRate)) // Sampling frequency
	binary.Write(file, binary.LittleEndian, uint32(sampleRate*1))
	binary.Write(file, binary.LittleEndian, uint16(1))
	binary.Write(file, binary.LittleEndian, uint16(8))
	binary.Write(file, binary.LittleEndian, []byte("data"))
	binary.Write(file, binary.LittleEndian, uint32(sampleRate*duration))

	// Генерация звука
	for t := 0; t < sampleRate*duration; t++ {
		binary.Write(file, binary.LittleEndian, uint8((t & int(t>>8))))
	}

	println("Done! output.wav generated.")
}
