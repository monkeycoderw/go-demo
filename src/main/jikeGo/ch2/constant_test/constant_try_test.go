package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

/**
常量定义
*/
func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry(t *testing.T) {

	a:=7 // 0111
	t.Log(Readable)
	t.Log(Writable)
	t.Log(Executable)
	a = a &^ Readable
	a = a &^ Writable
	t.Log(a&Readable==Readable, a&Writable==Writable, a&Executable==Executable)

}