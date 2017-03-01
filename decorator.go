package learn_pattern

import "fmt"

//手机抽象接口，即装饰者模式中的抽象组件类
type IPhone interface {
	Print()
}

type ApplePhone struct {
}

//苹果手机，即装饰着模式中的具体组件类
func (this *ApplePhone) Print() {
	fmt.Println("开始执行具体的对象——苹果手机")
}

// 贴膜，即具体装饰者
type Sticker struct {
	phone IPhone
}

func (this *Sticker) SetPhone(phone IPhone) {
	this.phone = phone
}

func (this *Sticker) Print() {
	this.phone.Print()
	this.AddSticker()
}

/// 新的行为方法
func (this *Sticker) AddSticker() {
	fmt.Println("现在苹果手机有贴膜了")
}

type Accessories struct {
	phone IPhone
}

func (this *Accessories) SetPhone(phone IPhone) {
	this.phone = phone
}

func (this *Accessories) Print() {
	this.phone.Print()
	this.AddAccessories()
}

/// 新的行为方法
func (this *Accessories) AddAccessories() {
	fmt.Println("现在苹果手机有漂亮的挂件了")
}
