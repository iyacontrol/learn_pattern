package learn_pattern

type Operation interface {
	GetResult() float64
	SetNumA(numA float64)
	SetNumB(numB float64)
}
type BaseOperation struct {
	numberA float64
	numberB float64
}

func (this *BaseOperation) SetNumA(numA float64) {
	this.numberA = numA
}

func (this *BaseOperation) SetNumB(numB float64) {
	this.numberB = numB
}

type OperationAdd struct {
	BaseOperation
}

func (this *OperationAdd) GetResult() float64 {
	return this.numberA + this.numberB
}

type OperationSub struct {
	BaseOperation
}

func (this *OperationSub) GetResult() float64 {
	return this.numberA - this.numberB
}

type OperationMul struct {
	BaseOperation
}

func (this *OperationMul) GetResult() float64 {
	return this.numberA * this.numberB
}

type OperationDiv struct {
	BaseOperation
}

func (this *OperationDiv) GetResult() float64 {
	if this.numberB == 0 {
		panic("被除数不能为0")
	}
	return this.numberA / this.numberB
}

type OperationFactory struct {
}

func (this *OperationFactory) CreateOperation(op string) (operation Operation) {
	switch op {
	case "+":
		operation = new(OperationAdd)
	case "-":
		operation = new(OperationSub)
	case "*":
		operation = new(OperationMul)
	case "/":
		operation = new(OperationDiv)
	default:
		panic("运算符号错误")
	}
	return
}
