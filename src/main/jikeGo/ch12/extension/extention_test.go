package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) speak() {
	fmt.Println("Pet Speak")
}

func (p *Pet) speakTo (host string) {
	p.speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	fmt.Println("Dog speak!")
}

func TestDog(t *testing.T) {
	dog := new (Dog)
	dog.speakTo("Nick")
}

