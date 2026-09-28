package main
import "fmt"

func main() {
	var r float64
	fmt.Print("Masukan jari jari lingkaran : ")
	fmt.Scan(&r)
	luas := 3.14159 * r * r
	fmt.Printf("Luas lingkaran adalah : %.1f\n", luas)
}
