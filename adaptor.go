package learn_pattern

import (
	"fmt"
)

type Target interface {
	Request()
}

type Adaptee struct{}

func (this *Adaptee) SpecificRequest() {
	fmt.Println("This is an special request !")
}

type Adaptor struct {
	adaptee Adaptee
}

func (this *Adaptor) Request() {
	this.adaptee.SpecificRequest()
}
