package main

// #include <stdint.h>
// uint64_t Add(uint64_t a, uint64_t b);
import "C"
import "fmt"

func main() {
	r := C.Add(12, 30)
	fmt.Println(r)
}
