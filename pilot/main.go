package main

import "fmt"

const AIRCRAFT1 = 1

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft Aircraft
}
type Aircraft int
