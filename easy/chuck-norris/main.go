package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Scan()

	var binaryBits string
	message := scanner.Text()

	for _, char := range message {
		binaryBits += fmt.Sprintf("%.7b", char)
	}

	var encodedString string
	prevBit := binaryBits[0]
	count := 1

	for i := 1; i < len(binaryBits); i++ {
		if binaryBits[i] == prevBit {
			count += 1
		} else {
			encodedString += encode(prevBit, count) + " "
			prevBit = binaryBits[i]
			count = 1
		}
	}

	encodedString += encode(prevBit, count)
	fmt.Println(encodedString)
}

func encode(bit byte, count int) string {
	result := "0"

	if bit == []byte("0")[0] {
		result += "0"
	}

	result += " " + strings.Repeat("0", count)
	return result
}
