package condition

import (
	"testing"
)

//func someFun() {
//
//}
//
//func TestMultiSelect(t *testing.T) {
//	if v,err:=someFun(); err==nil {
//		t.Log("OK",v);
//	} else {
//		t.Log("error")
//	}
//}

func TestSwitchMultiCase (t *testing.T) {

	for i:=0; i < 5; i++ {
		switch i {
		case 0,2:
			t.Log("even")
		case 1,3:
			t.Log("odd")
		default:
			t.Log("it is not 0-5")
		}
	}
	//switch os:= runtime.GOOS; os {
	//
	//}
}

/**

    if else的另一种写法

 */
func TestSwitchCondition(t *testing.T)  {
	for i:=0; i < 5; i++ {
		switch  {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unkonw")
		}
	}
}