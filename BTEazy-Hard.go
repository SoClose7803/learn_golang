package main

import (
	"fmt"
)
func sum(a,b int ) int {
	return a+b
}
func even_check(a int) string {
	if a%2==0{
		return "So chan"
	}
	return "So le"
}
func fatorial(a int) int {
	if a==0{
		return 1
	}
	return a*fatorial(a-1)
}

func main() {
	fmt.Println("Xin chào, thế giới!")
	fmt.Println("Tong cua 2 so 3 va 4 la:", sum(3,4))
	fmt.Println("So chan hay le:", even_check(3))
	fmt.Println("Giai thua cua 5 la:", fatorial(5))
} 