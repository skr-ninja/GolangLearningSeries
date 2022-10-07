package main

import (
	"fmt"
	"sort"
)

func main() {

	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	fmt.Println(a)
	var i int
	for i = 0; i < len(a); i++ {

		fmt.Println(i)
	}

	// array sorting
	b := []int{5, 3, 10, 23, 2, 12, 5, 6, 9}
	sort.Ints(b)
	fmt.Println(b)

	// sort array string
	//string := []string{"z", "b", "c", "j", "k", "n", "p"}
	//sort.Strings(string)
	//fmt.Println(string)

	// another example
	strArray := []string{"sunil", "Ram", "Awni", "Pankaj", "Abhimanyu", "Seeta", "Geeta"}
	sort.Strings(strArray)
	fmt.Println(strArray)

}
