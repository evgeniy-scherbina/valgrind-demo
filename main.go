package main

import "C"
import (
	"flag"
	"fmt"
	"time"
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
	delay := flag.Int("delay", 0, "")
	flag.Parse()

	fmt.Printf("delay start\n")
	time.Sleep(time.Duration(*delay) * time.Second)
	fmt.Printf("delay end\n")

	fmt.Printf("start\n")
	fmt.Printf("n: %v\n", *n)
	memoryLeak(*n)
	fmt.Printf("end\n")
}
