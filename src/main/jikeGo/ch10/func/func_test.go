package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
  多返回值
 */
func returnMultiValues () (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/**
  计时监控
 */
func timeSpent(inner func(op int) int) func (op int) int {
	return func(n int) int {
		start:= time.Now()
		ret:= inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

/**
  睡眠函数
 */
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func sum(ops ...int) int {
	ret := 0
	for _,op := range ops {
		ret += op
	}
	return ret
}

/**

 */
func TestFn(t *testing.T) {
	a,b := returnMultiValues()
	t.Log(a,b)
	tsSf := timeSpent(slowFun)
	t.Log(tsSf(10))
	t.Log(sum(1,2,3,4,5))
	t.Log(sum(1,2,3,4,5,6))

}

func clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
 	defer clear()
 	fmt.Println("Start")
 	panic("err")
}



