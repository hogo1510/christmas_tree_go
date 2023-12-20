package main
import "fmt"

func boom(height int) {
	for i := 1; i <= height; i++ {
		for j := 0; j < height-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	var hoogte int = 20
	boom(hoogte)
	for i := 2; i <= hoogte; i++{
		fmt.Print(" ")
	}
	fmt.Println("|| \n")
	fmt.Println("Fijne Kerst Iedereen ! ! !\n")
}
