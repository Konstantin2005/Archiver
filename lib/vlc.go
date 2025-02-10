package lib

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk
type BinaryChunk string

type HexChunks []HexChunk
type HexChunk string
type encodingTable map[rune]string

const ChunkSize = 8

func Encode(str string) string {
	str = strings.ToLower(str)
	bSTR := encodeBin(str)
	chunks := splitByChunks(bSTR, ChunkSize)

	fmt.Println("chunks:", chunks)
	return ""
}

func (hcs BinaryChunks) ToString() string {
	const sep = " "

	switch len(hcs) {
	case 0:

		return ""
	case 1:
		return string(hcs[0])
	}
	var duf strings.Builder

	duf.WriteString(string(hcs[0]))

	for _, chunk := range hcs[1:] {
		duf.WriteString(sep)
		duf.WriteString(string(chunk))
	}
	return duf.String()
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0)

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()

		res = append(res, hexChunk)
	}
	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, ChunkSize)
	if err != nil {
		panic("cannot convert chunk to hex" + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)

}

// prepareText prepare text to be fit for encode :
// changes upper case : 1 + lower case letter
// i.g: My name is Ted -> !my name is !Ted
func prepareText(str string) string {
	var duf strings.Builder
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			duf.WriteRune('!')
			duf.WriteRune(unicode.ToLower(ch))
		} else {
			duf.WriteRune(ch)
		}
	}
	return duf.String()
}

// разбивает binary string on binary shanks string with given size,
// i.g.: '100101010010101010101011' -> '10010101 00101010 10101011'
func splitByChunks(bStr string, ChunkSize int) BinaryChunks {

	strLen := utf8.RuneCountInString(bStr)

	ChunksCount := strLen / ChunkSize

	if strLen/ChunksCount != 0 {
		ChunksCount++
	}
	res := make(BinaryChunks, 0, ChunksCount)

	var duf strings.Builder

	for i, ch := range bStr {
		duf.WriteString(string(ch))

		if (i+1)%ChunkSize == 0 {
			res = append(res, BinaryChunk(duf.String()))
			duf.Reset()
		}
	}
	if duf.Len() != 0 {
		lastChunk := duf.String()

		lastChunk += strings.Repeat("0", ChunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}

func encodeBin(str string) string {
	var duf strings.Builder
	for _, ch := range str {
		duf.WriteString(bin(ch))
	}
	return duf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character %c" + string(ch))
	}
	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}
