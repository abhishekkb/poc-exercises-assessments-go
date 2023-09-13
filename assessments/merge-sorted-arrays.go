package assessments

import "fmt"

func MergeSortedArrays() {

	a := []int{}
	b := []int{}

	a = append(a, 2, 3, 5, 7, 8)
	b = append(b, 1, 2, 4, 6, 7, 9)

	la := len(a)
	lb := len(b)

	c := []int{}

	for i, j := 0, 0; i != la-1 && j != lb-1; {

		if a[i] < b[j] {
			fmt.Println("11111")
			c = append(c, a[i])
			i++
		} else if a[i] > b[j] {
			fmt.Println("22222")
			c = append(c, b[j])
			j++
		} else if a[i] == b[j] {
			fmt.Println("=======")
			c = append(c, a[i])
			i++
			j++
		}

		if i == la-1 && j < lb {
			fmt.Println("33333")
			c = append(c, b[j])
			j++
		} else if j == lb-1 && i < la {
			fmt.Println("444444")
			c = append(c, a[i])
			i++
		}
	}

	fmt.Printf("result %v\n", c)

}
