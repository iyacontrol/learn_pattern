package learn_pattern

import (
	"fmt"
	"sync"
)

type singleton struct {
	name string
}

func (this *singleton) SayName() {
	fmt.Println(this.name)
}

var single *singleton
var lock sync.Mutex = sync.Mutex{}

func GetSingleton(name string) *singleton {
	if single == nil {
		lock.Lock()

		if single == nil {
			single = &singleton{name: name}
		}
		lock.Unlock()
	}
	return single

}
