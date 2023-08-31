package a2

import (
	"fmt"
	"strconv"
)

func Somef() {
	t := []interface{}{"A", "B", "C", 1, 2, 3, "4", 5, "6"}

	a := []int{}
	b := []string{}
	for _, v := range t {
		switch v.(type) {
		case string:
			s := v.(string)
			if isInt(s) {
				i, _ := strconv.Atoi(s)
				a = append(a, i)
			} else {
				b = append(b, s)
			}
		case int:
			a = append(a, v.(int))
		}
	}
	fmt.Printf("int arr %v \n", a)
	fmt.Printf("string arr %v \n", b)
}

func isInt(s string) bool {
	_, e := strconv.Atoi(s)
	if e != nil {
		return false
	}
	return true
}
