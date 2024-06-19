package main

import (
    "github.com/01-edu/z01"
)

func ReverseBits(oct byte) byte {
    var result byte = 0
    for i := 0; i < 8; i++ {
   	 result <<= 1
   	 result |= oct & 1
   	 oct >>= 1
    }
    return result
}

func PrintByteAsBinary(b byte) {
    for i := 7; i >= 0; i-- {
   	 if b&(1<<i) != 0 {
   		 z01.PrintRune('1')
   	 } else {
   		 z01.PrintRune('0')
   	 }
    }
    z01.PrintRune('\n')
}

func main() {
    var testByte byte = 0x41 // 0010 0110 in binary
    reversedByte := ReverseBits(testByte)
    
    PrintByteAsBinary(testByte)
    PrintByteAsBinary(reversedByte)
}
