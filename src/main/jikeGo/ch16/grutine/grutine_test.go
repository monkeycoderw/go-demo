package grutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGrutine(t *testing.T) {
	for i := 0;i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}