package lzw

import (
	"fmt"
	"math"
)

//TODO comments
func Decompress(data []uint8) string{
	fmt.Print("Decompressing file... ")

	codes := toCodes(data)

	dict := initialDict()
	dictSize := int(math.Pow(2, 12))

	output := ""
	previous := ""
	nextKey := 256
	for i, code := range codes {
		if nextKey == dictSize {
			dict = initialDict()
			nextKey = 256
		}

		W, ok := dict[int(code)]
		if ok {
			if i != 0 {
				dict[nextKey] = ( previous + string(W[0]) )
				nextKey++
			}
			output += W

			previous = W
		} else {
			V := previous + string(previous[0]) // previous[0] will not fail as previous must be length > 1 for the map lookup to fail
			dict[nextKey] = V
			nextKey++
			output += V

			previous = V
		}
	}

	fmt.Println("Done.")
	return output
}

func toCodes(data []uint8) []uint16 {
	nBytes := len(data)
	codes := []uint16 {}
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

func initialDict() map[int]string {
	dict := map[int]string {}
	for i := 0; i < 256; i++ {
        dict[i] = string(rune(i))
    }

	return dict
}
