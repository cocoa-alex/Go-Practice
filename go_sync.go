package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex
var a string

func f() {
	a = "hello word"
	fmt.Println("解锁")
	l.Unlock()
	fmt.Println("解锁2")

}

func main() {
	l.Lock()
	fmt.Println("One")
	go f()
	fmt.Println("Two")
	l.Lock()
	fmt.Println("Three")
	fmt.Println(a)
}
