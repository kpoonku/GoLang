package main

import "fmt"

type person struct {
	firstName           string
	lastName            string
	favIceCreamFlavours []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	john := person{
		firstName:           "John",
		lastName:            "Copper",
		favIceCreamFlavours: []string{"vanila", "bubbleGum", "chocolate"},
	}

	will := person{
		firstName:           "Will",
		lastName:            "Smith",
		favIceCreamFlavours: []string{"pecan", "banana", "coconut"},
	}

	persons := []person{john, will}

	mapPerson := map[string]person{}

	for _, val := range persons {
		mapPerson[val.lastName] = val
	}

	fmt.Println(mapPerson)

	// loop the map
	for key, value := range mapPerson {
		fmt.Println(key, value)
	}

	sedanCar := sedan{
		vehicle: vehicle{doors: 4,
			color: "red"},
		luxury: true,
	}

	blackTruck := truck{
		vehicle: vehicle{doors: 4,
			color: "black"},
		fourWheel: true,
	}
	fmt.Println("blackTruck.color : ", blackTruck.color)
	fmt.Println("sedanCar.doors : ", sedanCar.doors)
	fmt.Println(sedanCar)
	fmt.Println(blackTruck)

	familySizeCar := struct {
		vehicle
		luxury bool
	}{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println("familySizeCar : ", familySizeCar)
}
