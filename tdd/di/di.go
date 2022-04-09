package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, msg string) {
    fmt.Fprintf(writer, "Hello, %s", msg)
}


func main(){
	Greet(os.Stdout, "Elodies")
}
