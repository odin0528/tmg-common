package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	tz, _ := time.LoadLocation("Asia/Taipei")

	Init(tz)

	c := 0
	AddJob(Every(1).Second(), func() {
		fmt.Println("c:",c)
		c++
	})

	b := 0
	AddJob(Every(100).Millisecond(), func() {
		fmt.Println("b:",b)
		b++
	})

	cron.Serve()

	time.Sleep(time.Second * 10)
}
