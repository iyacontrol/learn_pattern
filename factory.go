package learn_pattern

import "fmt"

//定义抽象产品
type ICar interface {
	Show()
}

//定义发动机
type Engine struct {
}

func (this *Engine) Show() {
	fmt.Println("发动机组装成功！")
}

//定义底盘
type Underpan struct {
}

func (this *Underpan) Show() {
	fmt.Println("底盘组装成功！")
}

//定义轮胎
type Wheel struct {
}

func (this *Wheel) Show() {
	fmt.Println("轮胎组装成功！")
}

//定义抽象工厂
type IFactory interface {
	CreateCar() ICar
}

type EngineFactory struct {
}

func (this *EngineFactory) CreateCar() ICar {
	return new(Engine)
}

type UnderpanFactory struct {
}

func (this *UnderpanFactory) CreateCar() ICar {
	return new(Underpan)
}

type WheelFactory struct {
}

func (this *WheelFactory) CreateCar() ICar {
	return new(Wheel)
}
