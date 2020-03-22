package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	//	encode("test.jpg", "encode.txt")
	decode("yas.txt", "decode_MS.txt")
}

func encode(read, write string) {
	fRead, _ := os.Open(read)
	defer fRead.Close()

	fWrite, _ := os.Create(write)
	defer fWrite.Close()

	const bufSize int = 1024
	readBuf := make([]byte, bufSize)
	encodeBuf := make([]byte, bufSize)

	var readError error
	var countRead int
	for readError == nil {
		countRead, readError = fRead.Read(readBuf)

		for i := 0; i <= countRead-1; i++ {
			encodeBuf[i] = ^readBuf[i]
		}

		strEncode := hex.EncodeToString(encodeBuf[:countRead])
		fWrite.WriteString(strEncode)
	}
}

func decode(read, write string) {
	fRead, _ := os.Open(read)
	defer fRead.Close()

	fWrite, _ := os.Create(write)
	defer fWrite.Close()

	const bufSize int = 1024
	readBuf := make([]byte, bufSize)
	decodeBuf := make([]byte, bufSize)
	invertBuf := make([]byte, bufSize)

	var readError, decodeError error
	var countRead, countDecode int
	for readError == nil {
		countRead, readError = fRead.Read(readBuf)

		countDecode, decodeError = hex.Decode(decodeBuf, readBuf[:countRead])
		for i := 0; i <= countDecode-1; i++ {
			invertBuf[i] = ^decodeBuf[i]
		}
		if decodeError != nil {
			fmt.Println("decode error:", decodeError)
		}
		fWrite.Write(invertBuf[:countDecode])

	}
}
