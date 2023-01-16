package lzw

import "fmt"

func Decompress(data []uint8) []uint8 {
	fmt.Print("Decompressing file... ")

	fmt.Println()
	codes := toCodes(data)
	fmt.Println(codes[:4])

	dict := map[int]string {}
	for i := 0; i < 256; i++ {
        dict[i] = string(rune(i))
    }
    fmt.Println(dict)

	fmt.Println("Done.")
	return nil
}

func toCodes(data []uint8) []uint16 {
	nBytes := len(data)
	codes := []uint16{}
	for i := 0; i < nBytes; i += 3 {
		if i < nBytes-2 {
			var byte1, byte2, byte3 uint8 = data[i], data[i+1], data[i+2]
			
			var code1 uint16 = (uint16(byte1) << 4) | (uint16(byte2) >> 4)
			var code2 uint16 = (uint16(byte2) & 0b00001111) << 8 | uint16(byte3)
			codes = append(codes, code1, code2)
		} else {
			if i == nBytes-1 {
				fmt.Println("ERROR: Last code only has 8 bits out of 12.")
				continue
			}

			var byte1, byte2 uint8 = data[i], data[i+1]
			
			var code uint16 = (uint16(byte1) << 8) | uint16(byte2)
			codes = append(codes, code)
		}
	}

	return codes
}
