package main

import (
	"fmt"
)

func main() {
	var i int = 0
	var j int = 0

	var chr = [13]uint{'C', 'E', 'A', 'D', 'Ш', 'S', 'A', 'Z', 'P', 'F', 'B', 'H', 'O'}

	fmt.Println("Mohammad Fatoni Anggris (1955617840-17)")
	fmt.Println("-----------------------------------------")

	for k := 0; k < 2; k++ {
		for l := 0; l <= len(chr); l++ {
			if i <= 4 && k == 0 {
				fmt.Printf("Nilai i = %v \n", i)
				i++
			}
			if l <= 10 && k == 1 {
				fmt.Printf("Nilai j = %v \n", j)
				if l == 5 {
					for m := 0; m <= len(chr); m++ {
						if m%2 == 0 {
							fmt.Printf("character = %U '%c' starts at byte position %v \n", chr[m], chr[m], m)
						}
					}
				}
				j++
			}
		}
	}

}
