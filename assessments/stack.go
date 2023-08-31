package assessments

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []string
}

func (s *Stack) pop() string {
	l := len(s.data) - 1
	p := s.data[l]
	s.data = s.data[:l]
	return p
}

func (s *Stack) push(d string) {
	s.data = append(s.data, d)
}

func (s *Stack) print() {
	fmt.Printf("Current Stack data --> %v \n", strings.Join(s.data, ","))
}

func (s *Stack) len() int {
	return len(s.data)
}

func DoStackOperations() {
	stack := Stack{[]string{}}

	stack.push("Hello1")
	stack.push("World1")
	stack.push("World2")
	stack.push("Hello2")
	stack.push("Hello3")
	stack.push("World3")
	stack.push("World4")
	stack.push("Hello4")

	for i := 0; stack.len() > 0; i++ {
		fmt.Println("================================================================")
		stack.print()
		fmt.Printf("Before popping data from stack, len = %d \n", stack.len())
		p := stack.pop()
		fmt.Printf("After popping data from stack, len = %d, popped val = %s \n", stack.len(), p)
		stack.print()
		fmt.Println("================================================================")
	}
}
