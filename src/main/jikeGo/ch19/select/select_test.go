package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("Working on OtherTask")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("OtherTask is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// buffer channel
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

/**
csp
*/
func TestSelect(t *testing.T) {
	select {
		case ret := <- AsyncService() :
			t.Log(ret)
		case <- time.After(time.Microsecond * 10):
			t.Error("time out")
	}
}