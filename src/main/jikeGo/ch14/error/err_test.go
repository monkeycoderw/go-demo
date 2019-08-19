package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should not be less than 2")
var MoreThan100Error = errors.New("n should be not large than 100")


func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, MoreThan100Error
	}
	fibList := []int{1,1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList,nil
}

func TestGetFibonacci(t *testing.T)  {
	if v, err := GetFibonacci(1); err!=nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
	} else {
		t.Log(v)
	}
}