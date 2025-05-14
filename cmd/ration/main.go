package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/B0TMirage/week-ration.git/internal/products"
)

func main() {
	fmt.Println("Приветствую в приложение \"Счастливый рацион\". Я твой персональный помощник.")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Рацион составляется на неделю по схеме 5-4-3-2-1.")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("5 овощей, 4 фрукта, 3 источника белка, 2 крупы, 1 сладость.")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("\nКогда будешь готов начать, нажми любую кнопку :З")
	fmt.Scanln()

	ration := products.GenerateRation()

	for {
		fmt.Println("Твой рацион готов: ")
		printRation(ration)

		fmt.Println("\nВозможные команды: ")
		var commandNumber int
		fmt.Println("1. Заменить продукт \n2. Полностью пересмотреть рацион")

	chooseComandLoop:
		for {
			fmt.Scanln(&commandNumber)

			switch commandNumber {
			case 1:
				var productNumber int

				fmt.Println("\nВведите номер продукта, который хотите заменить")
				for {
					fmt.Scanln(&productNumber)
					if productNumber < 1 || productNumber > 15 {
						fmt.Println("\nПродукта с данным номером нет в списке")
						continue
					}
					err := products.ReplaceProduct(&ration, productNumber)
					if err != nil {
						fmt.Println("\nНевозможно заменить продукт под данным номером, уникальные продукты данной категории закончились")
						time.Sleep(2 * time.Second)
					}
					break chooseComandLoop
				}
			case 2:
				ration = products.GenerateRation()
				break chooseComandLoop
			default:
				fmt.Println("Такой команды не существует")
				continue
			}
		}

		clearConsole()
	}

}

func printRation(r []string) {
	for i, v := range r {
		fmt.Printf("%d. %s\n", i+1, v)
		time.Sleep(50 * time.Millisecond)
	}
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Для Windows
	} else {
		cmd = exec.Command("clear") // Для Unix-подобных систем
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
