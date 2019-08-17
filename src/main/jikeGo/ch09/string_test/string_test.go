package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认值零值
	s = "hello"
	t.Log(len(s))
	// s[1] = '3' // string是不可变的byte slice
	s = "\xE4\xB8\xA5"
	// s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	s = "中"
	t.Log(len(s))
	c := []rune(s)
	// t.Log("rune size:", unsafe.Sizeof(c[0])
	t.Logf("中  unicode %x", c[0])
	t.Logf("中 utf8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s :="中华人民共和国"
	// 自动输出rune,不输出byte
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}