package assessments

import (
	"fmt"
	"sync"
)

func UseWaitGroups() {
	l := []int{}
	l = append(l, 2)
	l = append(l, 4)
	l = append(l, 21)
	l = append(l, 32)
	l = append(l, 9)

	l = append(l, 12)
	l = append(l, 44)
	l = append(l, 211)
	l = append(l, 320)
	l = append(l, 89)

	chunksLen := len(l) / 5
	wg := &sync.WaitGroup{}
	for i := 0; i < chunksLen; i++ {
		wg.Add(1)
		//fmt.Printf("chunk - %v\n", ll[i])
		c := i * 5
		chunk := l[c : c+5]
		fmt.Printf("chunk - %v\n", chunk)
		go func(l []int) {
			defer wg.Done()
			sum := 0
			for i := 0; i < len(l); i++ {
				sum += l[i]
			}
			fmt.Printf("sum %d\n", sum)
		}(chunk)
	}
	wg.Wait()
	fmt.Println("complete")
}
