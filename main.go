package main

import "C"
import (
	"flag"
	"fmt"
)

func memoryLeak(n int) {
	for i := 0; i < n; i++ {
		var uuid *C.uchar
		uuid = (*C.uchar)(C.malloc(16))
		_ = uuid
	}
}

func main() {
	n := flag.Int("n", 0, "")
	flag.Parse()

	fmt.Printf("start\n")
	fmt.Printf("n: %v\n", *n)
	memoryLeak(*n)
	fmt.Printf("end\n")
}
