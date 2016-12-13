package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/vbauerster/uiprogress"
)

const (
	totalItem    = 100
	maxBlockSize = 12
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	decor := func(s *uiprogress.Statistics) string {
		str := fmt.Sprintf("%d/%d", s.Completed, s.Total)
		return fmt.Sprintf("%-7s", str)
	}

	p := uiprogress.New()
	bar := p.AddBar(totalItem).AppendETA().PrependFunc(decor)

	blockSize := rand.Intn(maxBlockSize) + 1
	for i := 0; i < 100; i += 1 {
		time.Sleep(time.Duration(blockSize) * (50*time.Millisecond + time.Duration(rand.Intn(5*int(time.Millisecond)))))
		bar.Incr(1)
		blockSize = rand.Intn(maxBlockSize) + 1
	}

	p.WaitAndStop()
	fmt.Println("stop")
}