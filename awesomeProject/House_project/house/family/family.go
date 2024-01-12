package family

import "fmt"

type Person struct {
	Surname  string
	Name     string
	Age      int
	Birthday string
	Height   int
}

type Pet struct {
	PetType   string
	PetBreed  string
	Name      string
	Age       int
	HairColor string
}

func PrintFamily(x []interface{}) {
	for i := 0; i < len(x); i++ {
		switch v := x[i].(type) {

		//семья
		case Pet:
			fmt.Println(fmt.Sprintf("\tПитомец\n"+
				"Тип питомца:\t%s\n"+
				"Порода питомца:\t%s\n"+
				"Имя питомца:\t%s\n"+
				"Возраст питомца:\t%d\n"+
				"Цвет шерсти питомца:\t%s\n", v.PetType, v.PetBreed, v.Name, v.Age, v.HairColor))
		case Person:
			fmt.Println(fmt.Sprintf("\tСемья\n"+
				"Фамилия:\t%s\n"+
				"Имя:\t%s\n"+
				"Возраст:\t%d\n"+
				"День рождения:\t%s\n"+
				"Рост:\t%d\n", v.Surname, v.Name, v.Age, v.Birthday, v.Height))
		}
	}
}
