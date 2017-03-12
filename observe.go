package learn_pattern

import (
	"container/list"
	"fmt"
)

type Handler interface {
	Serve(int)
}

type HandlerFunc func(args int)

func (f HandlerFunc) Serve(args int) {
	f(args)
}

type Publish struct {
	HandlerFuncs *list.List
	Value        int
}

func NewPublisher() *Publish {
	p := new(Publish)
	p.HandlerFuncs = list.New()
	p.Value = 22
	return p
}

func (p *Publish) AddSubscriber(f HandlerFunc) { //可以添加任何方法，只要是HandlerFunc的
	p.HandlerFuncs.PushBack(f)
}

func (p *Publish) Notify(args int) {
	for f := p.HandlerFuncs.Front(); f != nil; f = f.Next() {
		f.Value.(Handler).Serve(args) //此处是我觉得比较得意的，比一般的观察者先进的地方，实现了任意函数名的调用
	}
}

type Subscribe struct {
}

func (s *Subscribe) Dosome(args int) {
	fmt.Println("Dosome:", args)
}
func (s *Subscribe) Fucksome(args int) {
	fmt.Println("Fucksome", args)
}
