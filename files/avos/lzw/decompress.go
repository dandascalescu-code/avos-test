package lzw

import (
	"fmt"
	"math"
)

/* Decodes the input data according to the LZW algorithm https://en.wikipedia.org/wiki/Lempel%E2%80%93Ziv%E2%80%93Welch#Decoding 
using 12-bit fixed-width encoding and the standard initial dictionary.
The dictionary is reset once full.
*/
func Decompress(data []uint8) string{
	fmt.Print("Decompressing file... ")

	codes := toCodes(data)

	dict := initialDict()
	dictSize := int(math.Pow(2, 12)) // max code in 12 bits (+1)

	output := ""
	previous := "" // the previous string emitted to the output
	nextKey := 256 // what position to input the next dictionary code (next available integer)
	for i, code := range codes {
		if nextKey == dictSize {
			dict = initialDict()
			nextKey = 256
		}

		W, ok := dict[int(code)]
		if ok {
			if i != 0 { // skip this on the first iter, as previous will be empty
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

/* Converts the data input into uint16s that contain the 12-bit codes.
Input data is an encoding of 12-bit codes packed into 8-bit bytes (such that 3 bytes contain 2 codes). 
The final code (if there are an odd number of them) is padded to fill the final 2 bytes of the input.
*/
func toCodes(data []uint8) []uint16 {
	nBytes := len(data)
	codes := []uint16 {}
	for i := 0; i < nBytes; i += 3 {
		if i < nBytes-2 {
			// Forms two 12-bit codes from each three bytes of the data
			var byte1, byte2, byte3 uint8 = data[i], data[i+1], data[i+2]
			
			var code1 uint16 = (uint16(byte1) << 4) | (uint16(byte2) >> 4)
			var code2 uint16 = (uint16(byte2) & 0b00001111) << 8 | uint16(byte3)
			codes = append(codes, code1, code2)
		} else { // Odd code
			if i == nBytes-1 {
				fmt.Println("ERROR: Last code only has 8 bits out of 12.")
				continue
			}
			// Forms the final 12-bit code from the final two bytes of the data
			var byte1, byte2 uint8 = data[i], data[i+1]
			fmt.Println(data[i], data[i+1])
			
			var code uint16 = (uint16(byte1) << 8) | uint16(byte2)
			//var code uint16 = (uint16(byte1) << 4) | (uint16(byte2) & 0b11110000) >> 4
			//var code uint16 = (uint16(byte1) & 0b00001111) << 8 | uint16(byte2)
			fmt.Println(code)
			codes = append(codes, code)
		}
	}

	return codes
}

/* Constructs the initial dictionary of codes 0-255 (the strings from "\x00" to "\xFF"), 
i.e. all matchable single-character strings
*/
func initialDict() map[int]string {
	dict := map[int]string {}
	for i := 0; i < 256; i++ {
        dict[i] = string(rune(i))
    }

	return dict
}
