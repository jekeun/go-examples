package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"time"
)

func main() {

	scheduler.Every(1).Seconds().Run(runStrtegy)

	// 무한루프, 단순 블록
	select {}
}

func runStrtegy() {
	fmt.Println(time.Now())
}