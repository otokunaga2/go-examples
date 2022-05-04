package main

import (
	"fmt"
	"io"
//	"time"
)
type Sleeper interface{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}
const countdownStart = 3
const finalWord = "Go!"
func Countdown(out io.Writer, spySleeper *SpySleeper) {
	for i := countdownStart; i > 0; i-- {
		spySleeper.Sleep()
		//time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)

	}
	//time.Sleep(1 * time.Second)

	spySleeper.Sleep()
	fmt.Fprint(out, finalWord)

}

func main() {
	//Countdown(os.Stdout)
}
