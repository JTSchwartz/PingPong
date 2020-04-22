/*******************************************************************************
/
/      filename:  PingPongFanInOrder.go
/
/   description:  Creates a fan of threads that alternates Ping and Pong
/
/        author:  Schwartz, Jacob
/      login id:  FA_18_CPS356_33
/
/         class:  CPS 356
/    instructor:  Schwartz
/    assignment:  Homework #6
/
/      assigned:  October 25, 2018
/           due:  November 1, 2018
/
/******************************************************************************/
package main

import (
	"fmt"
	"strconv"
	"sync"
)
import "os"

type Node struct {
	num   int
	ready chan int
}

func (c *Node) run(start <-chan int, wg *sync.WaitGroup) {
	<-start
	var pingpong = ""
	if c.num%2 == 0 {
		pingpong = "pong"
	} else {
		pingpong = "ping"
	}
	fmt.Println(pingpong, "// printed by goroutine", c.num)
	c.ready <- c.num
	wg.Done()
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	list := make([]Node, n)
	var wg sync.WaitGroup
	wg.Add(n)

	begin := make(chan int, 1)
	begin <- 1
	start := make(chan int, 1)
	node := Node{num: 1, ready: start}
	list[0] = node
	go node.run(begin, &wg)
	for i := 2; i <= n; i++ {
		last := list[i-2]
		fan := Node{num: i, ready: make(chan int, 1)}
		list[i-1] = fan
		go fan.run(last.ready, &wg)
	}
	wg.Wait()
}
