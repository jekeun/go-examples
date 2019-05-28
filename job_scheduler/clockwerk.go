package main

import (
	"fmt"
	"github.com/onatm/clockwerk"
	"time"
)

var count = 1

type DummyJob struct{}

func (d DummyJob) Run() {

	fmt.Printf("Every %d Seconds\n", count)
	count++
}

func main() {
	var job DummyJob
	c := clockwerk.New()

	c.Every(1 * time.Second).Do(job)
	c.Start()

	time.Sleep(100 * time.Second)
}
