package main

import "fmt"

func main() {
	maksimalAngka := 100

	bilPrima := make([]bool, maksimalAngka+1)

	for i := 2; i <= maksimalAngka; i++ {
		bilPrima[i] = true
	}

	for p := 2; p*p <= maksimalAngka; p++ {
		if bilPrima[p] {
			for i := p * p; i <= maksimalAngka; i += p {
				bilPrima[i] = false
			}
		}
	}

	fmt.Println("Angka Prima hingga", maksimalAngka, "adalah:")
	for i := 2; i <= maksimalAngka; i++ {
		if bilPrima[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	fmt.Println()

	var soal2 int
	fmt.Print("ini kelipatan 7 adalah \n")
	for soal2 = 0; soal2 < 100; soal2++ {
		if soal2 % 7 == 0 {
			fmt.Print(soal2, " ")
		}
	}
	fmt.Println()
	fmt.Println()
	
	var panjang int = 12
	var lebar int = 10
	var luas int = panjang * lebar
	fmt.Printf("%d x %d adalah : %v \n", panjang, lebar, luas)
}