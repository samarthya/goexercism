package main

import "sync"

import "fmt"

type myIntf interface {
	sayHello(wg *sync.WaitGroup)
}

type myStruct struct {
	id int
	sync.Mutex
}

func (ms *myStruct) sayHello(wg *sync.WaitGroup) {
	ms.Lock()
	defer ms.Unlock()
	defer wg.Done()
	fmt.Println(" Hello from : ", ms.id)
}

func main() {
	msa := make([]myStruct, 5)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		msa[i] = myStruct{id: i}
	}
	var t myIntf
	for _, v := range msa {
		t = v
		t.sayHello(&wg)
	}
	wg.Wait()
}
