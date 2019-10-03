package main

import "fmt"

type cpu6502 struct {
	a   byte
	x   byte
	y   byte
	pc  byte
	sp  byte
	mem [1 * 1024]byte
}

func newCPU6502() *cpu6502 {
	cpu := cpu6502{}
	return &cpu
}

func (cpu cpu6502) String() string {
	return fmt.Sprintf("A: %x X: %x Y: %x PC: %x SP: %x MEM: %x", cpu.a, cpu.x, cpu.y, cpu.pc, cpu.sp, cpu.mem)
}

func (cpu *cpu6502) sta(value byte) error {
	cpu.mem[value] = cpu.a
	return nil
}

func (cpu *cpu6502) lda(value byte) error {
	cpu.a = value
	return nil
}

func (cpu *cpu6502) ldx(value byte) error {
	cpu.x = value
	return nil
}

func (cpu *cpu6502) ldy(value byte) error {
	cpu.y = value
	return nil
}

func (cpu *cpu6502) inx() error {
	cpu.x++
	return nil
}

func (cpu *cpu6502) iny() error {
	cpu.y++
	return nil
}

func (cpu *cpu6502) txa() error {
	cpu.a = cpu.x
	return nil
}

func (cpu *cpu6502) tax() error {
	cpu.x = cpu.a
	return nil
}

func (cpu *cpu6502) tya() error {
	cpu.a = cpu.y
	return nil
}

func (cpu *cpu6502) tay() error {
	cpu.y = cpu.a
	return nil
}
