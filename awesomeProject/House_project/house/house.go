package house

import (
	"awesomeProject/House_project/house/devices"
	"awesomeProject/House_project/house/family"
	"awesomeProject/House_project/house/furniture"
	"awesomeProject/House_project/house/things"
	"fmt"
)

type House struct {
	Family    family.Person
	Devices   []devices.Devices
	Furniture []furniture.Furniture
	Things    []things.Things
	Rooms     int
	Square    int
}

func (h House) Print() {
	fmt.Print("Дом:\n")
	fmt.Print("Комнаты: ", h.Rooms, ", Площадь: ", h.Square, "\n\n")
	fmt.Print("\nТехника:\n")
	for _, tmpDevices := range h.Devices {
		tmpDevices.Print()
	}
	fmt.Print("\nМебель:\n")
	for _, tmpFurniture := range h.Furniture {
		tmpFurniture.Print()
	}
}

func MakeFamily() {
	var lfamily []interface{}
	Mom := family.Person{
		Surname:  "Соловьёва",
		Name:     "Алина",
		Age:      30,
		Birthday: "12.05.1993",
		Height:   167,
	}
	lfamily = append(lfamily, Mom)
	Dad := family.Person{
		Surname:  "Соловьёв",
		Name:     "Павел",
		Age:      31,
		Birthday: "18.10.1992",
		Height:   182,
	}
	lfamily = append(lfamily, Dad)
	Son := family.Person{
		Surname:  "Соловьёв",
		Name:     "Николай",
		Age:      8,
		Birthday: "25.10.2015",
		Height:   130,
	}
	lfamily = append(lfamily, Son)
	Daughter := family.Person{
		Surname:  "Соловьёва",
		Name:     "Кристина",
		Age:      5,
		Birthday: "1.08.2018",
		Height:   105,
	}
	lfamily = append(lfamily, Daughter)
	Cat := family.Pet{
		PetType:   "Кошка",
		PetBreed:  "Сиамская",
		Name:      "Муся",
		Age:       3,
		HairColor: "Белый",
	}
	lfamily = append(lfamily, Cat)
	family.PrintFamily(lfamily)
}
func Make() House {
	allFurniture := furniture.MakeFurniture()
	allDevices := devices.MakeDevices()
	allThings := things.MakeThings()

	house := House{
		Devices:   allDevices,
		Furniture: allFurniture,
		Things:    allThings,
		Rooms:     2,
		Square:    56,
	}
	return house
}
