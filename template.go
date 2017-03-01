package learn_pattern

import (
	"fmt"
)

type IPerson interface {
	SaySex()
}

type Person struct {
	Specific IPerson
}

func (this *Person) Introduce() {
	fmt.Println("Start Introduce!")
	if this.Specific != nil {
		this.Specific.SaySex()
	}
	fmt.Println("End Introduce!")
}

type Boy struct {
}

func (this *Boy) SaySex() {
	fmt.Println("I am a Boy!")
}

type Girl struct {
}

func (this *Girl) SaySex() {
	fmt.Println("I am a Girl!")
}
