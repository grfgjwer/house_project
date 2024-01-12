package furniture

import (
	"fmt"
)

type Furniture struct {
	Type     string
	Length   float32
	Width    float32
	Material string
	Color    string
}

func (f Furniture) Print() {
	fmt.Print("Тип мебели: ", f.Type, "\nДлина: ", f.Length, "\nШирина: ", f.Width, "\nМатериал: ", f.Material, "\nЦвет: ", f.Color)
}

func MakeFurniture() []Furniture {
	var furniture []Furniture
	sofa := Furniture{
		Type:     "Диван",
		Length:   180,
		Width:    80,
		Material: "Велюр",
		Color:    "Синий",
	}
	closet := Furniture{
		Type:     "Шкаф",
		Length:   200,
		Width:    70,
		Material: "Дерево",
		Color:    "Коричневый",
	}
	table := Furniture{
		Type:     "Стол",
		Length:   110,
		Width:    70,
		Material: "Дерево",
		Color:    "Черный",
	}
	chair := Furniture{
		Type:     "Стул",
		Length:   110,
		Width:    50,
		Material: "Металл",
		Color:    "Черный",
	}
	bed := Furniture{
		Type:     "Кровать",
		Length:   200,
		Width:    180,
		Material: "Металл",
		Color:    "Белый",
	}
	furniture = append(furniture, sofa, closet, table, chair, bed)
	return furniture
}
