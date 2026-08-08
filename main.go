package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "用法: crc <crc8|crc16|crc32|crc32c> <文本>")
		os.Exit(2)
	}
	v, err := Sum(os.Args[1], []byte(os.Args[2]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%#x\n", v)
}
