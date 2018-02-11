package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/datalayer/kuber/cmd"
)

func main() {
	go f()
	cmd.Execute()
}

func f() {
	for true {
		dowork(10)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func dowork(loopindex int) {
	// simulate work
	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("gr[%d]: %d\n", loopindex)
}
