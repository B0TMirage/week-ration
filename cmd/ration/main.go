package main

import (
	"fmt"
	"time"

	"github.com/B0TMirage/week-ration.git/internal/products"
)

func main() {
	ration := products.GenerateRation()
	printRation(ration)

}

func printRation(r []string) {
	for i, v := range r {
		fmt.Printf("%d. %s\n", i+1, v)
		time.Sleep(50 * time.Millisecond)
	}
}
