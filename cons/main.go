package main

import (
	"github.com/goutoujunshi/goserver/cons/cons"
	"sync"

)

func main(){
	 wg := sync.WaitGroup{}
	wg.Add(1)
	cons.Server()

	wg.Wait()
}
