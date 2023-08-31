package assessments

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Flight struct {
	Price int `json:"price"`
}

type Flights []Flight

func SortObjects() {
	flights := []Flight{}

	flights = append(flights, Flight{Price: 1000})
	flights = append(flights, Flight{Price: 19})
	flights = append(flights, Flight{Price: 40})
	flights = append(flights, Flight{Price: 190})

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})
	fmt.Println(flights)
	b, _ := json.Marshal(flights)
	fmt.Println(string(b))
}

func (fs Flights) Len() int {
	return len(fs)
}

func (fs Flights) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func (fs Flights) Less(i, j int) bool {
	return fs[i].Price < fs[j].Price
}

func SortUsingSortInterfaceMethod() {
	flights := []Flight{
		{Price: 10},
		{Price: 15},
		{Price: 110},
		{Price: 8},
		{Price: 6},
		{Price: 13},
	}
	sort.Sort(Flights(flights))
	fmt.Printf("Flights %v\n", flights)
}
