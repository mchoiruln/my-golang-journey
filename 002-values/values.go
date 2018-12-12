package main

import "fmt"

func SumString(a string, b string) string {
	return a + b
}

func SumInt(a int, b int) int {
	return a + b
}

func DivideInt(a int, b int) float64 {
	return float64(a / b)
}

func And(a bool, b bool) bool {
	return a && b
}

func Or(a bool, b bool) bool {
	return a || b
}

func Not(a bool) bool {
	return !a
}

func main() {
	fmt.Println(SumString("go", "lang"))
	fmt.Println(SumInt(1, 1))
	fmt.Println(DivideInt(24, 5))
	fmt.Println(And(true, false))
	fmt.Println(Or(true, false))
	fmt.Println(Not(true))
}
