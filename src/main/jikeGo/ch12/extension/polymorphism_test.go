package extension

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	writeHelloWorld() Code
}

type GoProgrammer struct {

}

func (p *GoProgrammer) writeHelloWorld() Code {
	return "fmt.Println(\"Go hello world\")"
}

type JavaProgrammer struct {

}

func (p *JavaProgrammer) writeHelloWorld() Code {
	return "System.out.println(\"Java hello world\")"
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.writeHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	//goProgrammer := &GoProgrammer{}
	goProgrammer := new(GoProgrammer)
	javaProgrammer := new(JavaProgrammer)

	writeFirstProgrammer(goProgrammer)
	writeFirstProgrammer(javaProgrammer)

}
