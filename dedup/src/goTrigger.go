package main

import "fmt"

func trigger(fun func(string) error) {
	defer wg.Done()
	for {
		fn, ok := <-pathChan
		if !ok {
			return
		}
		err := fun(fn)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done dedup file:", fn)
		}
	}

}
