package things

import (
	"fmt"
)

type Things struct {
	Type   string
	Length float32
	Width  float32
	Weight float32
	Color  string
}

func (d Things) Print() {
	fmt.Print("Вещь: ", d.Type, "\nДлина: ", d.Length, "\nШирина: ", d.Width, "\nВес: ", d.Weight, "\nЦветЖ", d.Color)
}

func MakeThings() []Things {
	var things []Things
	flower := Things{
		Type:   "Цветок",
		Length: 15,
		Width:  10,
		Weight: 0.4,
		Color:  "Зелёный",
	}
	painting := Things{
		Type:   "Картина",
		Length: 30,
		Width:  25,
		Weight: 0.2,
		Color:  "Разноцветный",
	}
	book := Things{
		Type:   "Книга",
		Length: 20,
		Width:  15,
		Weight: 0.2,
		Color:  "Черный",
	}
	mirror := Things{
		Type:   "Зеркало",
		Length: 15,
		Width:  10,
		Weight: 0.1,
		Color:  "Белый",
	}
	candle := Things{
		Type:   "Свеча",
		Length: 10,
		Width:  2,
		Weight: 0.1,
		Color:  "Белый",
	}
	things = append(things, flower, painting, book, mirror, candle)
	return things
}
