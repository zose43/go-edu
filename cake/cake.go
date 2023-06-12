package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose               bool
	Cakes                 int
	BakeTime              time.Duration
	BakeBaseDeviation     time.Duration
	BakeBuf               int
	NumIcers              int
	IceTime               time.Duration
	IceBaseDeviation      time.Duration
	IceBuf                int
	InscribeTime          time.Duration
	InscribeBaseDeviation time.Duration
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		count := cake(i)
		if s.Verbose {
			fmt.Println("baking", count)
		}
		work(s.BakeTime, s.BakeBaseDeviation)
		baked <- count
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for count := range baked {
		if s.Verbose {
			fmt.Println("icing", count)
		}
		work(s.IceTime, s.IceBaseDeviation)
		iced <- count
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		count := <-iced
		if s.Verbose {
			fmt.Println("inscribing", count)
		}
		work(s.InscribeTime, s.InscribeBaseDeviation)
		if s.Verbose {
			fmt.Println("finished", count)
		}
	}
}

func (s *Shop) Work(cycles int) {
	for i := 0; i < cycles; i++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
