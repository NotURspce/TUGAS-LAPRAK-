package main
import "fmt"

func main() {

	var fahrenheit float64
	
	// bagian struktur INPUT
	fmt.Print("Masukan suhu fahrenheit : ")
	fmt.Scan(&fahrenheit)

	//rumus konversi fahrn jadi celcius (F-32) x 5 / 9
	celcius := (fahrenheit - 32) * 5 / 9

	//struktur output
	fmt.Printf("Suhu celcius nya adalah : %.0f\n", celcius)
}
