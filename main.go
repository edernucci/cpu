package main

import "fmt"

func main() {
	cpu := newCPU6502()
	buffer := []byte("Hello world!")
	cpu.lda(buffer[0])
	cpu.sta(byte(0))
	cpu.lda(buffer[1])
	cpu.sta(byte(1))
	cpu.lda(buffer[2])
	cpu.sta(byte(2))

	fmt.Println(cpu)
}
