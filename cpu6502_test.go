package main

import "testing"

// var cpu *cpu6502
// var buffer []byte

// func TestMain(m *testing.M) {
// 	cpu = newCPU6502()
// 	buffer = []byte("Hello fucking world")
// }

func BenchmarkLDA(b *testing.B) {
	cpu := newCPU6502()
	buffer := []byte("Hello fucking world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.lda(buffer[5])
	}
}

func BenchmarkLDX(b *testing.B) {
	cpu := newCPU6502()
	buffer := []byte("Hello fucking world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.ldx(buffer[5])
	}
}

func BenchmarkLDY(b *testing.B) {
	cpu := newCPU6502()
	buffer := []byte("Hello fucking world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.ldy(buffer[5])
	}
}

func BenchmarkSTA(b *testing.B) {
	cpu := newCPU6502()
	buffer := []byte("Hello fucking world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.sta(buffer[5])
	}
}

func BenchmarkTXA(b *testing.B) {
	cpu := newCPU6502()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.txa()
	}
}

func BenchmarkINX(b *testing.B) {
	cpu := newCPU6502()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.inx()
	}
}

func BenchmarkINY(b *testing.B) {
	cpu := newCPU6502()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.iny()
	}
}

func BenchmarkTAX(b *testing.B) {
	cpu := newCPU6502()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.tax()
	}
}

func BenchmarkTYA(b *testing.B) {
	cpu := newCPU6502()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.tya()
	}
}

func BenchmarkTAY(b *testing.B) {
	cpu := newCPU6502()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpu.tay()
	}
}
