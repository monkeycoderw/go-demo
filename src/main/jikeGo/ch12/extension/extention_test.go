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
	pet *Pet
}

func (dog *Dog) speak() {
	//dog.pet.speak()
	fmt.Println("Dog speak!")
}

func (dog *Dog) speakTo(host string) {
	//dog.pet.speakTo(host)
	dog.speak()
	fmt.Println("", host)
}

func TestDog(t *testing.T) {
	dog := new (Dog)
	dog.speakTo("Nick")
}

