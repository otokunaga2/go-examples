package main

import (
	"fmt"
	"io"
	"time"
	"os"
)
type Sleeper interface{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type ConfigurableSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}
/*
type CountdownOpeartionSpy struct{
	Calls []string
}

func (s *CountdownOpeartionSpy) Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOpeartionSpy) Write()
*/

const countdownStart = 3
const finalWord = "Go!"
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		//time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)

	}


	sleeper.Sleep()
	fmt.Fprint(out, finalWord)

}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
