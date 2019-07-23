package archive

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for a := 0; a < 10; a++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * 4)
		}(a)
	}
	wg.Wait()
}
