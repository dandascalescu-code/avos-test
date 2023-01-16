package lzw

import "fmt"

func Decompress(data []uint8) []uint8 {
	fmt.Print("Decompressing file... ")

	fmt.Println()
	nBytes := len(data)
	fmt.Printf("%v bytes\n", nBytes)
	codes := []uint16{}
	for i := 0; i < nBytes; i += 3 {
		if i < nBytes-2 {
			var byte1, byte2, byte3 uint8 = data[i], data[i+1], data[i+2]
			
			var code1 uint16 = (uint16(byte1) << 4) | (uint16(byte2) >> 4)
			var code2 uint16 = (uint16(byte2) & 0b00001111) << 8 | uint16(byte3)
			codes = append(codes, code1, code2)
		} else {
			// TODO: handle edge case
		}
	}
	fmt.Println(codes[:4])

	fmt.Println("Done.")
	return nil
}
