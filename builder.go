package learn_pattern

import "container/list"
import "fmt"

type Product struct {
	parts *list.List
}

func (this *Product) Add(part string) {
	this.parts.PushBack(part)
}

func (this *Product) Show() {
	for e := this.parts.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

type IBuilder interface {
	BuildHead()
	BuildBody()
	BuildHand()
	BuildFeet()
	GetResult() Product
}

type FatPersonBuilder struct {
	prduct Product
}

func NewFatPerson() *FatPersonBuilder {
	l := list.New()
	p := Product{parts: l}
	f := &FatPersonBuilder{prduct: p}
	return f

}

func (this *FatPersonBuilder) BuildHead() {
	this.prduct.parts.PushBack("胖人头")
}

func (this *FatPersonBuilder) BuildBody() {
	this.prduct.parts.PushBack("胖人身")
}

func (this *FatPersonBuilder) BuildHand() {
	this.prduct.parts.PushBack("胖人手")
}

func (this *FatPersonBuilder) BuildFeet() {
	this.prduct.parts.PushBack("胖人脚")
}

func (this *FatPersonBuilder) GetResult() Product {
	return this.prduct
}

type ThinPersonBuilder struct {
	prduct Product
}

func NewThinPerson() *ThinPersonBuilder {
	l := list.New()
	p := Product{parts: l}
	f := &ThinPersonBuilder{prduct: p}
	return f

}

func (this *ThinPersonBuilder) BuildHead() {
	this.prduct.parts.PushBack("瘦人头")
}

func (this *ThinPersonBuilder) BuildBody() {
	this.prduct.parts.PushBack("瘦人身")
}

func (this *ThinPersonBuilder) BuildHand() {
	this.prduct.parts.PushBack("瘦人手")
}

func (this *ThinPersonBuilder) BuildFeet() {
	this.prduct.parts.PushBack("瘦人脚")
}

func (this *ThinPersonBuilder) GetResult() Product {
	return this.prduct
}

type Director struct {
}

func (this *Director) Construct(builder IBuilder) {
	builder.BuildHead()
	builder.BuildBody()
	builder.BuildHand()
	builder.BuildFeet()
}
