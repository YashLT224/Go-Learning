package Basic

import "fmt"


type Bike struct{
	Name string
	Colour string
	Price float64
}

type Vechile interface{
	ShowDetails()
	ShowName() string
}

func (bike Bike) ShowDetails(){               //
	fmt.Println("Bike Name:",bike.Name)
	fmt.Println("Bike Colour",bike.Colour)
	fmt.Println("Bike  price:",bike.Price)
}
func (bike Bike) ShowName()string{
	return bike.Name
}

func Interface(){
	var obj Vechile=Bike{
		Name:"Honda",
		Colour:"red",
		Price:56.5,
	}
	obj.ShowDetails()
	fmt.Println(obj.ShowName())
} 