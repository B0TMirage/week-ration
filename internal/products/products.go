package products

import (
	"math/rand"
	"time"
)

var vegetables, fruits, protein, cereal, sweets []string

func GenerateRation() []string {
	fillRation()
	availableRation := [][]string{vegetables, fruits, protein, cereal, sweets}
	readyRation := make([]string, 0, 15)

	for i, numberOfProducts := 0, 5; i < len(availableRation); i, numberOfProducts = i+1, numberOfProducts-1 {
		currentCategory := availableRation[i]
		for j := 0; j < numberOfProducts; j++ {
			randVegIndex := rand.Intn(len(currentCategory))
			readyRation = append(readyRation, currentCategory[randVegIndex])
			time.Sleep(10 * time.Millisecond)
			currentCategory = append(currentCategory[:randVegIndex], currentCategory[randVegIndex+1:]...)
		}
	}

	return readyRation
}

func fillRation() {
	vegetables = []string{"огурец", "помидоры", "баклажаны", "кабачки", "тыква", "болгарский перец", "брокколи", "цветная капуста", "стручковая фасоль", "морковь", "шпинат", "репчатый лук", "красный лук", "брюссельская капуста", "зелёный горошек", "кукуруза", "спаржа", "редис", "капуста белокочанная", "капуста китайская", "картофель", "краснокачанная капуста", "свекла", "зелёный лук", "укроп", "петрушка", "розмарин", "сельдерей", "кинза", "грибы", "салат айсберг", "авокадо", "батат", "руколу", "цукини", "лук порей"}
	fruits = []string{"яблоко", "банан", "апельсин", "мандарин", "киви", "груша", "слива", "виноград зелёный", "виноград красный", "голубика", "клубника", "вишня", "малина", "персики", "нектарины"}
	protein = []string{"курица", "говядина", "свинина", "кролик", "тунец"}
	cereal = []string{"рис", "гречка", "булгур"}
	sweets = []string{"белый шоколад"}
}
