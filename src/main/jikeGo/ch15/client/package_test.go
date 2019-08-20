package client

import "testing"
import "main/jikeGo/ch15/series"

func TestPackage(t *testing.T) {
	ret := series.GetFibonacciSeries(10)
	t.Log(ret)
}
