package type_test

import "testing"

type MyInt int64

/**
 *
 * 不支持隐式类型转换
 *
 */
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64

	b = int64(a)

	var c MyInt

	// cannot use b (type int64) as type MyInt in assignment
	// c = b
	c = MyInt(b)
	t.Log(a,b,c)

}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a

	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// string是一个零值，空字符串
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("S is nil")
	}
}
