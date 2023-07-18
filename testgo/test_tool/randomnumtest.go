package test_tool

import (
	"fmt"
	"math/rand"
	"time"
)

func Randomnumtest() {
	p := fmt.Print
	pl := fmt.Println
	pl("random number test start")
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	pl()

	pl(rand.Float64())

	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	pl()

	seed1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed1)
	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	pl()
	seed2 := rand.NewSource(42)
	r2 := rand.New(seed2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	pl()
	seed3 := rand.NewSource(42)
	r3 := rand.New(seed3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))
	pl()

	pl("random number test end")
}
