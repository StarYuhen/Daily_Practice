package main

import "time"

func main() {
	ch := make(chan int)
	go func() {
		time.AfterFunc(time.Second*60, func() {

		})
	}()

	go func() {
		for true {
			select {
			case <-ch:
				return

			}
		}

	}()
}
