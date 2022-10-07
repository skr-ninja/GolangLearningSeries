package main

import (
	"fmt"
	"math"
)

func main() {

	var name string = "Sunil kumar Rathour"
	println(name)

	data := "This is go lang"
	println(data)

	var number int = 100
	println(number)

	num := 50
	println(num)

	var x float32 = 24.34
	fmt.Println(x)

	// You can declare multiple variables at once.
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	fmt.Println(1 + 2)

	fmt.Println("5.3/2.3=0", 5.3/2.3)

	fmt.Println(true && false)

	// conversions

	var i, j int = 3, 4

	println(i, j)

	var f float64 = math.Sqrt(float64(i*i + j*j))

	println(f)

	var z uint = uint(f)
	fmt.Println(i, j, z)

}
