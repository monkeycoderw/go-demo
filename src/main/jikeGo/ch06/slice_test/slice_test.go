package slice_test

import "testing"

/**
定义切片
 */
func TestSliceInit(t *testing.T) {

	// slice 没有长度定义
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1:=[]int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	s2:=make([]int,3,5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])

}

/**
切片可长
 */
func TestSliceGrouing(t *testing.T) {
	s := []int{}
	for i:=0;i<10;i++ {
		s = append(s,i)
		t.Log(len(s), cap(s))
	}
}

/**
	切片的内存共享
 */
func TestSliceShareMomery(t *testing.T) {
	year := []string{"jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer:=year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t testing.T) {
	a := []int{1,2,3,4}
	b := []int{1,2,3,4}
	// 切片不能进行比较
	a == b
}