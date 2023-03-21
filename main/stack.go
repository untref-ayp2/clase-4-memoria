package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	f1()
}

func f1() {
	f2()
}

func f2() {
	f3()
}

func f3() {
	f4()
}

func f4() {
	// debug.PrintStack()
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}
