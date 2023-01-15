package main

// Basic Threading With Go Routines

import (
	"runtime"
	"time"
)

func main() {
	// basit bir go routine kullanımı
	runtime.GOMAXPROCS(8)
	ch := make(chan string)
	go xFunc(ch)
	go printer(ch)
	time.Sleep(100 * time.Millisecond)
}

func xFunc(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
}

func printer(ch chan string) {
	for {
		println(<-ch)
	}
}
