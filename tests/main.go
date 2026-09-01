package main

import (
	"fmt"
)

const (
	PI          = 3.14159265358979323846
	ANTIGRAVITY = 9.80665
)

type Age int
type Salary float64

type Employee struct {
	Name   string
	Age    Age
	Salary Salary
}

func main() {
	// Math(10)
	// MapFunc()
	Age(30).PrintEmployeeDetails(Salary(50000.0))
	employee := Employee{
		Name:   "John Doe",
		Age:    30,
		Salary: 50000.0,
	}
	employee.PrintDetails()
}

func (e Employee) PrintDetails() {
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Salary:", e.Salary)
}

func (a Age) PrintEmployeeDetails(s Salary) {
	fmt.Println("Age:", a)
	fmt.Println("Salary:", s)

}

func MapFunc() {
	mymap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, value := range mymap {
		defer fmt.Println(key, value)
	}

	var value int = 2
	switch value {
	case 1:
		fmt.Println("Value is one")
	case 2:
		fmt.Println("Value is two")
		fallthrough
	case 3:
		fmt.Println("Value is three")
	default:
		fmt.Println("Value is unknown")
	}
}

func Math(number int) {
	for i := 0; i < number; i++ {
		fmt.Println("Math is fun!")
	}

	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
}

func ArraySlice() (int, string) {

	mmap := make(chan int)

	go func() {
		mmap <- 52
	}()

	lengt := len(mmap)

	fmt.Println(<-mmap)
	fmt.Println(lengt)

	return <-mmap, "done"
}
