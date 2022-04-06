package main

import (
	"github.com/jonasmateus/LoadBalancerGo"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)

	lb := LoadBalancerGo.New("Entry", "3000")

	wg.Add(1)
	go func(){

		lb.Serve()
	}()

	wg.Add(1)
	go func(){

		lb.ServeBackend("SERVER (1)", "3031")
	}()

	wg.Add(1)
	go func(){

		lb.ServeBackend("SERVER (2)", "3032")
	}()

	wg.Add(1)
	go func(){

		lb.ServeBackend("SERVER (3)", "3033")
	}()

	wg.Add(1)
	go func(){

		lb.ServeBackend("SERVER (4)", "3034")
	}()
	
	wg.Wait()
}