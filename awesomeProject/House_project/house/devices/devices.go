package devices

import (
	"fmt"
)

type Devices struct {
	Type   string
	Length float32
	Width  float32
	Weight float32
	Color  string
}

func (d Devices) Print() {
	fmt.Print("Тип устройства: ", d.Type, "\nДлина: ", d.Length, "\nШирина: ", d.Width, "\nВес: ", d.Weight, "\nЦветЖ", d.Color)
}

func MakeDevices() []Devices {
	var devices []Devices
	laptop := Devices{
		Type:   "Ноутбук",
		Length: 25,
		Width:  37,
		Weight: 2,
		Color:  "Черный",
	}
	phone := Devices{
		Type:   "Телефон",
		Length: 15,
		Width:  7,
		Weight: 0.2,
		Color:  "Черный",
	}
	TV := Devices{
		Type:   "Телевизор",
		Length: 64,
		Width:  113,
		Weight: 10,
		Color:  "Черный",
	}
	watch := Devices{
		Type:   "Часы",
		Length: 4,
		Width:  1,
		Weight: 0.5,
		Color:  "Черный",
	}
	fridge := Devices{
		Type:   "Холодильник",
		Length: 180,
		Width:  60,
		Weight: 70,
		Color:  "Белый",
	}
	devices = append(devices, laptop, phone, TV, watch, fridge)
	return devices
}
