package products

import (
	"errors"
	"fmt"
	"math/rand"
)

var vegetables, fruits, protein, cereal, sweets []string

func GenerateRation() []string {
	fillRation()
	availableRation := [][]string{vegetables, fruits, protein, cereal, sweets}
	readyRation := make([]string, 0, 15)

	for i, numberOfProducts := 0, 5; i < len(availableRation); i, numberOfProducts = i+1, numberOfProducts-1 {
		currentCategory := availableRation[i]
		for j := 0; j < numberOfProducts; j++ {
			readyRation = append(readyRation, getProduct(&currentCategory))
		}
	}

	return readyRation
}

func ReplaceProduct(ration *[]string, index int) error {
	if index <= 5 {
		if len(vegetables) <= 1 {
			return errors.New("not enough objects in the array")
		}
		(*ration)[index-1] = getProduct(&vegetables)
	} else if index <= 9 {
		if len(fruits) <= 1 {
			return errors.New("not enough objects in the array")
		}
		(*ration)[index-1] = getProduct(&fruits)
	} else if index <= 12 {
		if len(protein) <= 1 {
			return errors.New("not enough objects in the array")
		}
		(*ration)[index-1] = getProduct(&protein)
	} else if index <= 14 {
		if len(cereal) <= 1 {
			return errors.New("not enough objects in the array")
		}
		(*ration)[index-1] = getProduct(&cereal)
	} else {
		fmt.Println(len(sweets))
		if len(sweets) <= 1 {
			return errors.New("not enough objects in the array")
		}
		(*ration)[index-1] = getProduct(&sweets)
	}

	return nil
}

func getProduct(p *[]string) string {
	i := rand.Intn(len((*p)))
	product := (*p)[i]
	(*p) = append((*p)[:i], (*p)[i+1:]...)

	return product
}

func fillRation() {
	vegetables = []string{"огурец", "помидоры", "баклажаны", "кабачки", "тыква", "болгарский перец", "брокколи", "цветная капуста", "стручковая фасоль", "морковь", "шпинат", "репчатый лук", "красный лук", "брюссельская капуста", "зелёный горошек", "кукуруза", "спаржа", "редис", "капуста белокочанная", "капуста китайская", "картофель", "краснокачанная капуста", "свекла", "зелёный лук", "укроп", "петрушка", "розмарин", "сельдерей", "кинза", "грибы", "салат айсберг", "авокадо", "батат", "руколу", "цукини", "лук порей", "редька", "репа", "патисоны", "фасоль белая", "фасоль красная", "ламинария", "щавель"}
	fruits = []string{"яблоко", "банан", "апельсин", "мандарин", "киви", "груша", "слива", "виноград зелёный", "виноград красный", "голубика", "клубника", "вишня", "малина", "персики", "нектарины", "питахайя", "фейхоа", "манго", "папайя", "арбуз", "дыня", "черешня"}
	protein = []string{"курица", "говядина", "свинина", "кролик", "тунец"}
	cereal = []string{"рис", "гречка", "булгур"}
	sweets = []string{"белый шоколад"}
}
