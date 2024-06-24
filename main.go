package main

import "fmt"

func main() {
	var input float64

	fmt.Println("Geben Sie eine Länge (in Metern) ein: ")

	_, err := fmt.Scan(&input)
	if err != nil {
		return
	}

	fmt.Println(input, "Meter entsprechen: ")
	fmt.Println(input*1000, "mm")
	fmt.Println(input/1000, "km")
	fmt.Printf("%.3f Zoll \n", input*100/2.54)
	fmt.Printf("%.3g sm \n", input/1852)
	fmt.Printf("%.3g Lj", input/9460730472580800)

}
