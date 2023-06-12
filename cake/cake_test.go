package cake

import (
	"testing"
	"time"
)

var defaults = Shop{
	Cakes:        20,
	NumIcers:     1,
	BakeTime:     10 * time.Millisecond,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

func BenchmarkDefault(b *testing.B) {
	// no buffers, no base deviation
	shop := defaults
	shop.Work(b.N) // 229 milliseconds
}

func BenchmarkBuffers(b *testing.B) {
	//  add buffers, no base deviation
	shop := defaults
	shop.BakeBuf = 10
	shop.IceBuf = 10
	shop.Work(b.N) // 230 milliseconds
}

func BenchmarkDeviate(b *testing.B) {
	//  add deviation, no buffers
	shop := defaults
	shop.BakeBaseDeviation = shop.BakeTime / 4
	shop.IceBaseDeviation = shop.IceTime / 4
	shop.InscribeBaseDeviation = shop.InscribeTime / 4
	shop.Work(b.N) // 273 milliseconds, memory increase
}

func BenchmarkDeviateWithBuffers(b *testing.B) {
	//  add deviation, add buffers
	shop := defaults
	shop.BakeBaseDeviation = shop.BakeTime / 4
	shop.IceBaseDeviation = shop.IceTime / 4
	shop.InscribeBaseDeviation = shop.InscribeTime / 4
	shop.IceBuf = 10
	shop.BakeBuf = 10
	shop.Work(b.N) // 250 milliseconds
}

func BenchmarkMiddleSlower(b *testing.B) {
	//  making the middle stage slower
	shop := defaults
	shop.IceTime = 50 * time.Millisecond
	shop.Work(b.N) // 1036 milliseconds, memory and alloc *3 bigger
}

func BenchmarkBeginSlower(b *testing.B) {
	//  making the beginning stage slower
	shop := defaults
	shop.BakeTime = 50 * time.Millisecond
	shop.Work(b.N) // 1034 milliseconds, memory and alloc normally
}

func BenchmarkMiddleSlowerParallel(b *testing.B) {
	//  making the middle stage slower with parallels
	shop := defaults
	shop.IceTime = 50 * time.Millisecond
	shop.NumIcers = 5
	shop.Work(b.N) // 275 milliseconds, memory and alloc *2 bigger
}
