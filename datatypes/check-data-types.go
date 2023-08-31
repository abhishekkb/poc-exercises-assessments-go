package datatypes

import "fmt"

func CheckDataTypes() {
	checkAndPrintType(100)
	checkAndPrintType(`
		asdfjasdfalsdhfasdfjk
		sdlkfajsdflaksjdflkjs
	`)
	checkAndPrintType(map[string]int{})

	typeAssertion()
}

func checkAndPrintType(v interface{}) {
	switch t := v.(type) {
	case int:
		println("data given is of int type")
	case string:
		println("data given is of string type")
	default:
		fmt.Printf("data given is of unknown type %v \n", t)
	}
}

func typeAssertion() {
	fmt.Println("Doing type assertions")
	var x interface{} = "some val"

	str, ok := x.(string)

	if ok {
		fmt.Printf("is string type, val = %s \n", str)
	} else {
		fmt.Println("not a string")
	}
}
