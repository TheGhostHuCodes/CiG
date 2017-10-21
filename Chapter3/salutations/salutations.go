package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"Hello", "Greetings", "Good Day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
